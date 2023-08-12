package sTrading

import (
	"fmt"

	"github.com/yasseldg/simplego/sStr"
)

// States

type State int

const (
	State_Win         = State(1)
	State_Loss        = State(-1)
	State_InTrade     = State(2)
	State_CloseTs     = State(3) // CloseTs: Max In Trade Time permitted from Entry
	State_CancelTs    = State(4) // CancelTs: Max Excursion Time permitted from Trigger
	State_CancelPrice = State(5) // CancelPrice: Max Favorable Excursion permitted from Entry
	State_NotEntry    = State(6) // NotEntry: Never reached Entry Price
	State_DEFAULT     = State(0)
)

func GetState(s string) State {
	switch sStr.Lower(s) {
	case "win", "1":
		return State_Win
	case "loss", "-1":
		return State_Loss
	case "intrade", "in_trade", "2":
		return State_InTrade
	case "closets", "close_ts", "3":
		return State_CloseTs
	case "cancelts", "cancel_ts", "4":
		return State_CancelTs
	case "cancelprice", "cancel_price", "5":
		return State_CancelPrice
	case "notentry", "not_entry", "6":
		return State_NotEntry

	default:
		return State_DEFAULT
	}
}

func (s State) String() string {
	switch s {
	case State_Win:
		return "Win"
	case State_Loss:
		return "Loss"
	case State_InTrade:
		return "InTrade"
	case State_CloseTs:
		return "CloseTs"
	case State_CancelTs:
		return "CancelTs"
	case State_CancelPrice:
		return "CancelPrice"
	case State_NotEntry:
		return "NotEntry"
	default:
		return "-"
	}
}

func (s State) Log() string {
	return fmt.Sprintf(" ( %s ) ", s.String())
}
