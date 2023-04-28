package sCandle

func GetIntervalMinutes(interval string) int64 {
	switch interval {
	case "M1", "candle1m", "1m", "1", "1min":
		return 1
	case "M5", "candle5m", "5m", "5", "5min":
		return 5
	case "M15", "candle15m", "15m", "15", "15min":
		return 15
	case "H1", "candle1H", "1H", "60", "1h":
		return 60
	case "H4", "candle4H", "4H", "240", "4h":
		return 240
	case "D", "candle1D", "1Dutc", "1440", "1d":
		return 1440
	case "W":
		return 10080
	default:
		return 0
	}
}

func GetIntervalSeconds(interval string) int64 {
	return GetIntervalMinutes(interval) * 60
}

// IsClosing, ts is in seconds
func IsClosing(ts int64, interval string) bool {
	return (ts % GetIntervalSeconds(interval)) == 0
}

func GetIntervalSecondsMilli(interval string) int64 {
	return GetIntervalSeconds(interval) * 1000
}

// IsClosingMilli, ts is in milliseconds
func IsClosingMilli(ts int64, interval string) bool {
	return IsClosing(ts/1000, interval)
}

// PrevTs, ts is in seconds
func PrevTs(ts int64, interval string) int64 {
	intSec := GetIntervalSeconds(interval)
	diff := ts % intSec
	return ts - diff
}

// NextTs, ts is in seconds
func NextTs(ts int64, interval string) int64 {
	intSec := GetIntervalSeconds(interval)
	diff := ts % intSec
	return ts - diff + intSec
}

func GetIntervalGraphql(interval string) string {
	switch interval {
	case "candle1m", "1m":
		return "M1"
	case "candle5m", "5m":
		return "M5"
	case "candle15m", "15m":
		return "M15"
	case "candle1H", "1H":
		return "H1"
	case "candle4H", "4H":
		return "H4"
	case "candle1D", "1Dutc":
		return "D"
	default:
		return ""
	}
}
