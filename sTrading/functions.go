package sTrading

import "github.com/yasseldg/simplego/sConv"

// TakeProfitStoploss: returns take_profit and stop_loss prices
func TakeProfitStoploss(entry_price, take_profit_perc, stop_loss_perc float64, side Side) (take_profit, stop_loss float64) {
	switch side {
	case Side_Buy:
		take_profit = sConv.GetWithPercent(entry_price, take_profit_perc, true)
		stop_loss = sConv.GetWithPercent(entry_price, stop_loss_perc, false)

	case Side_Sell:
		take_profit = sConv.GetWithPercent(entry_price, take_profit_perc, false)
		stop_loss = sConv.GetWithPercent(entry_price, stop_loss_perc, true)
	}
	return
}
