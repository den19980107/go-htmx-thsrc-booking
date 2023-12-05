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

func (s StationCode) Code() int {
	switch s {
	case Nangang:
		return 1
	case Taipei:
		return 2
	case Banqiao:
		return 3
	case Taoyuan:
		return 4
	case Hsinchu:
		return 5
	case Miaoli:
		return 6
	case Taichung:
		return 7
	case Changhua:
		return 8
	case Yunlin:
		return 9
	case Chiayi:
		return 10
	case Tainan:
		return 11
	case Zuouing:
		return 12
	default:
		return 0
	}
}
