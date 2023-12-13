package crawler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/station"
)

const (
	BASE_URL           = "https://irs.thsrc.com.tw"
	BOOKING_PAGE_URL   = "https://irs.thsrc.com.tw/IMINT/?locale=tw"
	SUBMIT_FORM_URL    = "https://irs.thsrc.com.tw/IMINT/;jsessionid=%s?wicket:interface=:0:BookingS1Form::IFormSubmitListener"
	CONFIRM_TRAIN_URL  = "https://irs.thsrc.com.tw/IMINT/?wicket:interface=:1:BookingS2Form::IFormSubmitListener"
	CONFIRM_TICKET_URL = "https://irs.thsrc.com.tw/IMINT/?wicket:interface=:2:BookingS3Form::IFormSubmitListener"

	USER_AGENT      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"
	ACCEPT_HTML     = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	ACCEPT_IMG      = "image/webp,*/*"
	ACCEPT_LANGUAGE = "zh-TW,zh;q=0.8,en-US;q=0.5,en;q=0.3"
	ACCEPT_ENCODING = "gzip, deflate, br"

	BOOKING_PAGE_HOST = "irs.thsrc.com.tw"
)

type Crawler struct {
	collector *colly.Collector
}

func Create() Crawler {
	return Crawler{
		collector: colly.NewCollector(
			colly.AllowURLRevisit(),
			colly.MaxDepth(5),
			colly.UserAgent(USER_AGENT),
		),
	}
}

func (c *Crawler) SetCookies(url string, cookies []*http.Cookie) error {
	return c.collector.SetCookies(url, cookies)
}

func (c *Crawler) CompleteOrder(order ent.Order, captcha string, jsessionId string) error {
	trainDatas, feedBackErrors, err := c.submitForm(order, captcha, jsessionId)
	if err != nil {
		return fmt.Errorf("submit train order failed, err: %s", err)
	}

	if len(feedBackErrors) > 0 {
		errorStr := ""
		for index, err := range feedBackErrors {
			errorStr += err.Error()
			if index != len(feedBackErrors)-1 {
				errorStr += "\n"
			}
		}
		return errors.New(errorStr)
	}

	if len(trainDatas) == 0 {
		return errors.New("no train avaliable")
	}

	err = c.confirmTrain(trainDatas[0])
	if err != nil {
		return fmt.Errorf("confirm train failed, err: %s", err)
	}

	err = c.confirmTicket(order)
	if err != nil {
		return fmt.Errorf("confirm ticket failed, err: %s", err)
	}

	return nil
}

func (c *Crawler) GetCaptchaImageAndJsessionId(orderId int) (imageBase64 string, jsessionId string, cookieStr string, err error) {
	var wg sync.WaitGroup
	wg.Add(1)

	collector := cloneCollectorWithBasicEvent(c.collector, "get captcha code and jession id")
	collector.OnHTML("img[id='BookingS1Form_homeCaptcha_passCode']", func(e *colly.HTMLElement) {
		siteCookies := collector.Cookies(e.Request.URL.String())
		for _, cookie := range siteCookies {
			if cookie.Name == "JSESSIONID" {
				jsessionId = cookie.Value
			}
		}

		imgPath := e.Attr("src")
		url := BASE_URL + imgPath
		imageBase64, err = imageUrlToBase64(url)
		if err != nil {
			err = fmt.Errorf("convert image to base64 failed, err: %s", err)
		}

		wg.Done()
	})

	collector.Visit(BOOKING_PAGE_URL)
	wg.Wait()

	cookies := collector.Cookies(BOOKING_PAGE_URL)
	for _, c := range cookies {
		cookieStr = cookieStr + fmt.Sprintf("%s=%s; ", c.Name, c.Value)
	}

	return imageBase64, jsessionId, cookieStr, err
}

func (c *Crawler) submitForm(order ent.Order, captchaResult string, jsessionId string) ([]TrainData, []FeedBackError, error) {
	toTimeTable := formatTimeForTHSRC(order.StartTime)
	submitForm := SubmitForm{
		TripConTypesoftrip:            "0",
		TrainConTrainRadioGroup:       "0",
		SeatConSeatRadioGroup:         "0",
		BookingMethod:                 "radio31",
		SelectStartStation:            fmt.Sprintf("%d", station.StationCode(order.DepartureStation).Code()),
		SelectDestinationStation:      fmt.Sprintf("%d", station.StationCode(order.ArrivalStation).Code()),
		ToTimeInputField:              order.StartTime.Format("2006/01/02"),
		BackTimeInputField:            order.EndTime.Format("2006/01/02"),
		ToTimeTable:                   toTimeTable,
		TicketPanelRows0TicketAmount:  "1F",
		TicketPanelRows1TicketAmount:  "0H",
		TicketPannelRows2TicketAmount: "0W",
		TicketPanelRows3TicketAmount:  "0E",
		TicketPanelRows4TicketAmount:  "0P",
		HomeCaptchaSecurityCode:       captchaResult,
	}
	log.Printf("%+v", submitForm)
	submitJsonStr, _ := json.Marshal(submitForm)
	submitBody := make(map[string]string)
	json.Unmarshal(submitJsonStr, &submitBody)

	wg := sync.WaitGroup{}
	wg.Add(1)

	trainDatas := []TrainData{}
	feedBackErrors := []FeedBackError{}

	collector := cloneCollectorWithBasicEvent(c.collector, "submit form")

	// detect predict captcha failed
	collector.OnHTML("li[class='feedbackPanelERROR']", func(h *colly.HTMLElement) {
		h.ForEach("span", func(i int, h *colly.HTMLElement) {
			feedBackErrors = append(feedBackErrors, GetFeedBackError(h.Text))
		})
	})

	// if predict succes, should get response of list of train
	collector.OnHTML("div[class='result-listing']", func(h *colly.HTMLElement) {
		h.ForEach("div[class='btn-radio']", func(i int, h *colly.HTMLElement) {
			trainDatas = append(trainDatas, TrainData{
				TrainCode:     h.ChildAttr("input[name='TrainQueryDataViewPanel:TrainGroup']", "QueryCode"),
				DepartureTime: h.ChildAttr("input[name='TrainQueryDataViewPanel:TrainGroup']", "QueryDeparture"),
				ArrivalTime:   h.ChildAttr("input[name='TrainQueryDataViewPanel:TrainGroup']", "QueryArrival"),
				Date:          h.ChildAttr("input[name='TrainQueryDataViewPanel:TrainGroup']", "QueryDepartureDate"),
				Value:         h.ChildAttr("input[name='TrainQueryDataViewPanel:TrainGroup']", "value"),
			})
		})
	})

	collector.OnScraped(func(r *colly.Response) {
		defer wg.Done()
	})

	err := collector.Post(fmt.Sprintf(SUBMIT_FORM_URL, jsessionId), submitBody)

	wg.Wait()

	if feedBackErrors != nil {
		return nil, feedBackErrors, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return trainDatas, nil, nil
}

func (c *Crawler) confirmTrain(trainData TrainData) error {
	collector := cloneCollectorWithBasicEvent(c.collector, "confirm train")
	confirmTrainForm := ConfirmTrain{
		BookingS2FormHf0:                  "",
		TrainQueryDataViewPanelTrainGroup: trainData.Value,
	}

	confirmTrainJsonStr, _ := json.Marshal(confirmTrainForm)
	confirmTrainBody := make(map[string]string)
	err := json.Unmarshal(confirmTrainJsonStr, &confirmTrainBody)
	if err != nil {
		return fmt.Errorf("unmarshal confirm train json string failed, err: %s", err)
	}

	return collector.Post(CONFIRM_TRAIN_URL, confirmTrainBody)
}

func (c *Crawler) confirmTicket(order ent.Order) error {
	collector := cloneCollectorWithBasicEvent(c.collector, "confirm ticket")
	confirmTicketForm := ConfirmTicket{
		BookingS3FormSpHf0: "",
		DiffOver:           "1",
		IsSPromotion:       "1",
		PassengerCount:     "1",
		IsGoBackM:          "",
		BackHome:           "",
		TgoError:           "1",
		IdInputRadio:       "0",
		DummyId:            order.IDNumber,
		DummyPhone:         order.PhoneNumber,
		Email:              order.Email,
		Agree:              "on",
		TicketMemberSystemInputPanelTakerMemberSystemDataViewMemberSystemRadioGroup: "radio44",
	}
	confirmTicketJsonStr, _ := json.Marshal(confirmTicketForm)
	confirmTicketBody := make(map[string]string)
	err := json.Unmarshal(confirmTicketJsonStr, &confirmTicketBody)
	if err != nil {
		return fmt.Errorf("unmarshal confirm ticket json string failed, err: %s", err)
	}

	return collector.Post(CONFIRM_TICKET_URL, confirmTicketBody)
}

func cloneCollectorWithBasicEvent(collector *colly.Collector, execName string) *colly.Collector {
	clone := collector.Clone()
	clone.OnRequest(func(r *colly.Request) {
		r.Headers.Add("HOST", BOOKING_PAGE_HOST)
		r.Headers.Add("Accept", ACCEPT_HTML)
		r.Headers.Add("Accept-Language", ACCEPT_LANGUAGE)
		r.Headers.Add("Accept-Encoding", ACCEPT_ENCODING)
		r.Headers.Add("Connection", "keep-alive")
		r.Headers.Add("Content-Type", "application/x-www-form-urlencoded")
		log.Println("Visiting", r.URL)
	})

	clone.OnError(func(r *colly.Response, err error) {
		log.Printf("%s", err.Error())
	})

	clone.OnResponse(func(r *colly.Response) {
		log.Printf("%s on response with status code: %d", execName, r.StatusCode)
		if r.StatusCode == 200 {
			_ = os.WriteFile(fmt.Sprintf("%s.html", execName), r.Body, 0644)
		}
	})

	return clone
}

func imageUrlToBase64(URL string) (string, error) {
	//Get the response bytes from the url
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Accept", ACCEPT_HTML)
	req.Header.Set("Accept-Language", ACCEPT_LANGUAGE)
	req.Header.Set("Accept-Encoding", ACCEPT_ENCODING)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return "", err
	}
	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedStr, nil
}

func formatTimeForTHSRC(t time.Time) string {
	departureTime := t.Local().Format("03:04 PM")
	parts := strings.Split(departureTime, " ")
	hourMin := parts[0]
	meridiem := parts[1]
	timePart := strings.Split(hourMin, ":")
	hour, _ := strconv.Atoi(timePart[0])
	min, _ := strconv.Atoi(timePart[1])
	minStr := ""
	if min > 30 {
		minStr = "00"
	} else {
		minStr = "30"
	}
	if meridiem == "AM" {
		meridiem = "A"
	} else {
		meridiem = "P"
	}
	return fmt.Sprintf("%d%s%s", hour, minStr, meridiem)
}
