package sCandle

import (
	"strings"

	"github.com/yasseldg/simplego/sConv"
)

func GetIntervalMinutes(interval string) int64 {
	min := int64(0)
	interval, mult := getIntervalMult(interval)
	switch GetInterval(interval) {
	case Interval_1m:
		min = 1
	case Interval_5m:
		min = 5
	case Interval_15m:
		min = 15
	case Interval_1h:
		min = 60
	case Interval_4h:
		min = 240
	case Interval_D:
		min = 1440
	case Interval_W:
		min = 10080
	default:
		min = 0
	}
	return min * int64(mult)
}

func getIntervalMult(interval string) (string, int) {
	mult := 1
	strs := strings.Split(interval, "*")
	if len(strs) > 1 {
		interval = strs[0]
		mult_2 := sConv.GetInt(strs[1])
		if mult_2 > mult {
			mult = mult_2
		}
	}
	return interval, mult
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
