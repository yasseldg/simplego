package sCandle

import (
	"math"

	"github.com/yasseldg/simplego/sStat"
)

func (c Candle) LogReturn() float64 {
	if c.Close == 0 || c.Open == 0 {
		return 0
	}
	return sStat.ValidFloat64(math.Log(c.Close / c.Open))
}
