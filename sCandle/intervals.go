package sCandle

type Interval string

const (
	Interval_1m  = Interval("1m")
	Interval_5m  = Interval("5m")
	Interval_15m = Interval("15m")
	Interval_1h  = Interval("1h")
	Interval_4h  = Interval("4h")
	Interval_D   = Interval("D")
	Interval_W   = Interval("W")

	Interval_DEFAULT = Interval("DEFAULT")
)

func GetInterval(interval string) Interval {
	switch interval {
	case "1m":
		return Interval_1m
	case "5m":
		return Interval_5m
	case "15m":
		return Interval_15m
	case "1h":
		return Interval_1h
	case "4h":
		return Interval_4h
	case "D":
		return Interval_D
	case "W":
		return Interval_W
	default:
		return Interval_DEFAULT
	}
}
