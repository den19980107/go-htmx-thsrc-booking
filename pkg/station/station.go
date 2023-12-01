package station

type StationCode string

const (
	Nangang  StationCode = "Nangang"
	Taipei   StationCode = "Taipei"
	Banqiao  StationCode = "Banqiao"
	Taoyuan  StationCode = "Taoyuan"
	Hsinchu  StationCode = "Hsinchu"
	Miaoli   StationCode = "Miaoli"
	Taichung StationCode = "Taichung"
	Changhua StationCode = "Changhua"
	Yunlin   StationCode = "Yunlin"
	Chiayi   StationCode = "Chiayi"
	Tainan   StationCode = "Tainan"
	Zuouing  StationCode = "Zuouing"
)

func GetAllStations() []StationCode {
	return []StationCode{
		Nangang,
		Taipei,
		Banqiao,
		Taoyuan,
		Hsinchu,
		Miaoli,
		Taichung,
		Changhua,
		Yunlin,
		Chiayi,
		Tainan,
		Zuouing,
	}
}

// func (s StationCode) String() string {
// 	switch s {
// 	case Nangang:
// 		return "南港"
// 	case Taipei:
// 		return "台北"
// 	case Banqiao:
// 		return "板橋"
// 	case Taoyuan:
// 		return "桃園"
// 	case Hsinchu:
// 		return "新竹"
// 	case Miaoli:
// 		return "苗栗"
// 	case Taichung:
// 		return "台中"
// 	case Changhua:
// 		return "彰化"
// 	case Yunlin:
// 		return "雲林"
// 	case Chiayi:
// 		return "嘉義"
// 	case Tainan:
// 		return "台南"
// 	case Zuouing:
// 		return "左營"
// 	default:
// 		return "未知"
// 	}
// }
