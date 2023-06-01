package sDate

import (
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	ticker := NewTicker(10, nil)

	if ticker.Inc != 10 {
		t.Errorf("Unexpected increment value: got %v, want 10", ticker.Inc)
	}

	if ticker.Func == nil {
		t.Error("Function should not be nil")
	}
}

func TestTicker_Start(t *testing.T) {
	ts := time.Now().Unix()
	ticker := NewTicker(10, nil)
	ticker.Start(ts)

	if ticker.Ts != ts {
		t.Errorf("Unexpected timestamp: got %v, want %v", ticker.Ts, ts)
	}
}

func TestTicker_Update(t *testing.T) {
	ts := time.Now().Unix()
	ticker := NewTicker(10, nil)
	ticker.Start(ts)
	newTs := ts + 1000
	ticker.Update(newTs)

	if ticker.Ts != newTs {
		t.Errorf("Unexpected timestamp after update: got %v, want %v", ticker.Ts, newTs)
	}
}
