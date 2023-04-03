package sTrading

import (
	"github.com/yasseldg/simplego/sStr"
)

func GetSide(s string) int {
	switch sStr.Lower(s) {
	case "buy", "1":
		return SideBuy
	case "sell", "0":
		return SideSell
	default:
		return -1
	}
}

func IsBuy(s string) bool {
	return GetSide(s) == SideBuy
}

func IsSell(s string) bool {
	return GetSide(s) == SideSell
}

func GetSideStr(s int) string {
	switch s {
	case SideBuy:
		return "Buy"
	case SideSell:
		return "Sell"
	default:
		return ""
	}
}
