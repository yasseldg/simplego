package sTrading

import (
	"fmt"

	"github.com/yasseldg/simplego/sStr"
)

// *** Sides

type Side int

const (
	Side_Buy     = Side(1)
	Side_Sell    = Side(-1)
	Side_DEFAULT = Side(0)
)

func GetSide(s string) Side {
	switch sStr.Lower(s) {
	case "buy", "long", "1":
		return Side_Buy
	case "sell", "short", "-1":
		return Side_Sell
	default:
		return Side_DEFAULT
	}
}

func (s Side) IsBuy() bool {
	return s == Side_Buy
}

func (s Side) IsSell() bool {
	return s == Side_Sell
}

func (s Side) String() string {
	switch s {
	case Side_Buy:
		return "Buy"
	case Side_Sell:
		return "Sell"
	default:
		return ""
	}
}

func (s Side) Position() string {
	switch s {
	case Side_Buy:
		return "Long"
	case Side_Sell:
		return "Short"
	default:
		return ""
	}
}

func (s Side) Log() string {
	return fmt.Sprintf(" ( %s ) ", s.String())
}

func (s Side) LogPosition() string {
	return fmt.Sprintf(" ( %s ) ", s.Position())
}
