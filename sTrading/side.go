package sTrading

import (
	"github.com/yasseldg/simplego/sStr"
)

const (
	// Side
	Buy  = 1
	Sell = 2

	BuyStr  = "Buy"
	SellStr = "Sell"
)

func GetSide(s string) int {
	switch sStr.Lower(s) {
	case "buy", "1":
		return Buy
	case "sell", "2":
		return Sell
	default:
		return 0
	}
}

func IsBuy(s string) bool {
	return GetSide(s) == Buy
}

func IsSell(s string) bool {
	return GetSide(s) == Sell
}

func GetSideStr(s int) string {
	switch s {
	case Buy:
		return BuyStr
	case Sell:
		return SellStr
	default:
		return ""
	}
}
