package sCandle

import (
	"fmt"
	"math"

	"github.com/yasseldg/simplego/sDate"
	"github.com/yasseldg/simplego/sStat"
)

func (c Candle) LogReturn() float64 {
	if c.Close == 0 || c.Open == 0 {
		return 0
	}
	return sStat.ValidFloat64(math.Log(c.Close / c.Open))
}

func (o OHLC) String(prec int) string {
	return fmt.Sprintf("o: %*f  ..  h: %*f  ..  l: %*f  ..  c: %*f", prec, o.Open, prec, o.High, prec, o.Low, prec, o.Close)
}

func (o OHLCV) String(prec int) string {
	return fmt.Sprintf("%s  ..  v: %*f", o.OHLC.String(prec), prec, o.Volume)
}

func (c Candle) String(prec int) string {
	return fmt.Sprintf("Candle: %s  ..  %s", sDate.ForLog(c.UnixTs, 0), c.OHLCV.String(prec))
}
