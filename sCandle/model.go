package sCandle

type OHLC struct {
	Open  float64 `bson:"o" json:"o"`
	High  float64 `bson:"h" json:"h"`
	Low   float64 `bson:"l" json:"l"`
	Close float64 `bson:"c" json:"c"`
}

type Candle struct {
	UnixTs int64 `bson:"ts" json:"ts"`
	OHLC   `bson:",inline"`
}

type Candles []*Candle
