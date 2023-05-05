package sDate

import (
	"time"

	"github.com/yasseldg/simplego/sLog"
)

type Ticker struct {
	Ts   int64
	Inc  int64
	T    *time.Ticker
	Func TsFunc
}

type TsFunc func(ts int64)

func NewTicker(inc int64, f TsFunc) *Ticker {
	if f == nil {
		f = defaultFunc
	}
	t := new(Ticker)
	t.Inc = inc
	t.T = new(time.Ticker)
	t.Func = f
	return t
}

func (t *Ticker) Start(ts int64) {
	t.T = time.NewTicker(time.Duration(t.Inc) * time.Second)
	t.Update(ts)
	go t.loop()
}

func (t *Ticker) Update(ts int64) {
	if ts != t.Ts {
		t.Ts = ts
		t.T.Reset(time.Duration(t.Inc) * time.Second)
	}
}

func (t *Ticker) loop() {
	for {
		<-t.T.C
		t.Ts += t.Inc
		t.Func(t.Ts)
	}
}

func defaultFunc(ts int64) {
	sLog.Debug("ts: %s", ForLog(ts, 0))
}
