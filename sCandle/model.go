package sCandle

type OHLC struct {
	Open  float64 `bson:"o" json:"o"`
	High  float64 `bson:"h" json:"h"`
	Low   float64 `bson:"l" json:"l"`
	Close float64 `bson:"c" json:"c"`
}
type OHLCs []*OHLC

type OHLCV struct {
	OHLC   `bson:",inline"`
	Volume float64 `bson:"v" json:"v"`
}
type OHLCVs []*OHLCV

type Candle struct {
	UnixTs int64 `bson:"ts" json:"ts"`
	OHLCV  `bson:",inline"`
}
type Candles []*Candle
