package sCandle

func VerifyZeros(obj *OHLC) bool {
	return obj.Open <= 0 || obj.High <= 0 || obj.Low <= 0 || obj.Close <= 0
}

func VerifyInterval(obj_1, obj_2 *Candle, interval int64) bool {
	return (obj_2.UnixTs - obj_1.UnixTs) != interval
}

func VerifyOpenClose(obj_1, obj_2 *Candle) bool {
	return obj_1.Close != obj_2.Open
}
