package sDate

import (
	"time"
)

type Ticker struct {
	Ts  int64
	Inc int64
	T   *time.Ticker
	C   chan int64
}

func NewTicker(inc int64) *Ticker {
	t := new(Ticker)
	t.Inc = inc
	t.T = new(time.Ticker)
	t.C = make(chan int64)
	return t
}

func (t *Ticker) Start(ts int64) {
	t.T = time.NewTicker(time.Duration(t.Inc) * time.Second)
	t.Update(ts)
	go t.tickerLoop()
	go t.loop()
}

func (t *Ticker) Update(ts int64) {
	if ts != t.Ts {
		t.Ts = ts
		t.T.Reset(time.Duration(t.Inc) * time.Second)
	}
}

func (t *Ticker) tickerLoop() {
	for {
		<-t.T.C
		t.Ts += t.Inc
		t.C <- t.Ts
	}
}

func (t *Ticker) loop() {
	for {
		<-t.C
	}
}
