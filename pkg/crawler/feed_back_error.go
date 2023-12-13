package crawler

import (
	"fmt"
	"strings"
)

type FeedBackError string

const (
	Unknow                    FeedBackError = "未知錯誤: %s"
	CaptchaNotCorrect         FeedBackError = "檢測碼輸入錯誤"
	DipartureDateNotAvaliable FeedBackError = "去程您所選擇的日期超過目前開放預訂之日期"
)

func GetFeedBackError(text string) FeedBackError {
	if strings.Contains(text, string(CaptchaNotCorrect)) {
		return CaptchaNotCorrect
	}

	if strings.Contains(text, string(DipartureDateNotAvaliable)) {
		return DipartureDateNotAvaliable
	}

	return FeedBackError(fmt.Sprintf(string(Unknow), text))
}
