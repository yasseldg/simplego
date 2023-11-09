package sCandle

type Interval string

const (
	Interval_1m  = Interval("1m")
	Interval_5m  = Interval("5m")
	Interval_15m = Interval("15m")
	Interval_30m = Interval("30m")
	Interval_1h  = Interval("1h")
	Interval_4h  = Interval("4h")
	Interval_D   = Interval("D")
	Interval_W   = Interval("W")
	Interval_M   = Interval("M")

	Interval_DEFAULT = Interval("DEFAULT")
)

func (i Interval) IsDefault() bool {
	return i == Interval_DEFAULT
}

func (i Interval) String() string {
	if i.IsDefault() {
		return ""
	}
	return string(i)
}

func (i Interval) Seconds() int64 {
	if i.IsDefault() {
		return 0
	}
	return GetIntervalSeconds(i.String())
}

func (i Interval) Minutes() int64 {
	if i.IsDefault() {
		return 0
	}
	return GetIntervalMinutes(i.String())
}

func (i Interval) SecondsMilli() int64 {
	if i.IsDefault() {
		return 0
	}
	return GetIntervalSecondsMilli(i.String())
}

func (i Interval) IsClosing(ts int64) bool {
	if i.IsDefault() {
		return false
	}
	return IsClosing(ts, i.String())
}

func (i Interval) IsClosingMilli(ts int64) bool {
	if i.IsDefault() {
		return false
	}
	return IsClosingMilli(ts, i.String())
}

func (i Interval) Prev(ts int64) int64 {
	if i.IsDefault() {
		return 0
	}
	return PrevTs(ts, i.String())
}

func (i Interval) Next(ts int64) int64 {
	if i.IsDefault() {
		return 0
	}
	return NextTs(ts, i.String())
}

func GetInterval(interval string) Interval {
	switch interval {
	case "M1", "candle1m", "1m", "1", "1min":
		return Interval_1m
	case "M5", "candle5m", "5m", "5", "5min":
		return Interval_5m
	case "M15", "candle15m", "15m", "15", "15min":
		return Interval_15m
	case "M30", "candle30m", "30m", "30", "30min":
		return Interval_30m
	case "H1", "candle1H", "1H", "60", "1h":
		return Interval_1h
	case "H4", "candle4H", "4H", "240", "4h":
		return Interval_4h
	case "D", "candle1D", "1D", "1Dutc", "1440", "1d":
		return Interval_D
	case "W", "candle1W", "1W", "Week", "10080", "1w":
		return Interval_W
	case "M", "Month":
		return Interval_M
	default:
		return Interval_DEFAULT
	}
}
