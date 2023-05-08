package sCandle

import (
	"fmt"

	"github.com/yasseldg/simplego/sDate"
)

// GetFixes, prev is the previous candle, candles is the list of candles to be fixed, interval is the candle interval in seconds
// return de candles from next to prev to the last candle in candles
func GetFixes(prev *Candle, candles Candles, interval int64) Candles {
	news := Candles{}
	for _, candle := range candles {
		if prev.UnixTs >= candle.UnixTs {
			continue
		}
		if !VerifyInterval(prev, candle, interval) {
			news = append(news, candlesFromClose(prev.UnixTs, candle.UnixTs, interval, prev.Close)...)
		}
		fixCandle(candle, prev.Close)
		news = append(news, candle)
		prev = candle
	}
	return news
}

func Verify(candles Candles, interval int64) []error {
	errs := []error{}
	var prev *Candle
	for _, candle := range candles {
		if prev == nil {
			prev = candle
			continue
		}
		err := VerifyOne(prev, candle, interval)
		if err != nil {
			errs = append(errs, err)
		}
		prev = candle
	}
	return errs
}

func VerifyOne(prevObj, obj *Candle, interval int64) error {
	errBan := false
	log := fmt.Sprintf("ts: %d | %s ", prevObj.UnixTs, sDate.FormatD(prevObj.UnixTs, 2))

	if !VerifyInterval(prevObj, obj, interval) {
		log = fmt.Sprintf("%s |  %s != %s ", log, sDate.FormatD(prevObj.UnixTs, 2), sDate.FormatD(obj.UnixTs, 2))
		errBan = true
	}

	if !VerifyOpenClose(prevObj, obj) {
		log = fmt.Sprintf("%s |  close: %f != open: %f ", log, prevObj.Close, obj.Open)
		errBan = true
	}

	if !VerifyZeros(prevObj) {
		log = fmt.Sprintf("%s |  open: %f - high: %f - low: %f - close: %f ", log, prevObj.Open, prevObj.High, prevObj.Low, prevObj.Close)
		errBan = true
	}

	if !VerifyHighLow(prevObj) {
		log = fmt.Sprintf("%s |  low: %f > high: %f ", log, prevObj.Low, prevObj.High)
		errBan = true
	} else {
		if !VerifyLows(prevObj) {
			log = fmt.Sprintf("%s |  low: %f > open: %f - close: %f ", log, prevObj.Low, prevObj.Open, prevObj.Close)
			errBan = true
		}

		if !VerifyHighs(prevObj) {
			log = fmt.Sprintf("%s |  high: %f < open: %f - close: %f ", log, prevObj.High, prevObj.Open, prevObj.Close)
			errBan = true
		}
	}

	if errBan {
		return fmt.Errorf("%s", log)
	}
	return nil
}

func VerifyInterval(obj_1, obj_2 *Candle, interval int64) bool {
	return (obj_2.UnixTs - obj_1.UnixTs) == interval
}

func VerifyOpenClose(obj_1, obj_2 *Candle) bool {
	return obj_1.Close == obj_2.Open
}

func VerifyZeros(obj *Candle) bool {
	return obj.Open > 0 && obj.High > 0 && obj.Low > 0 && obj.Close > 0
}

func VerifyHighLow(obj *Candle) bool {
	return obj.High >= obj.Low
}

func VerifyHighs(obj *Candle) bool {
	return obj.High >= obj.Open && obj.High >= obj.Close
}

func VerifyLows(obj *Candle) bool {
	return obj.Low <= obj.Open && obj.Low <= obj.Close
}

func candlesFromClose(tsFrom, tsTo, interval int64, price float64) Candles {
	tsFrom += interval
	news := Candles{}
	for tsFrom < tsTo {
		news = append(news, &Candle{OHLCV: OHLCV{OHLC: OHLC{Open: price, High: price, Low: price, Close: price}, Volume: 0}, UnixTs: tsFrom})
		tsFrom += interval
	}
	return news
}

func fixCandle(candle *Candle, prevClosePrice float64) {
	candle.Open = prevClosePrice
	if candle.High < candle.Open {
		candle.High = candle.Open
	}
	if candle.Low > candle.Open || candle.Low == 0 {
		candle.Low = candle.Open
	}
	if candle.Close > candle.High {
		candle.High = candle.Close
	}
	if candle.Close < candle.Low && candle.Close > 0 {
		candle.Low = candle.Close
	}
	if candle.Close == 0 {
		candle.Close = candle.Open
	}
}
