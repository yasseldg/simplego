package sTrading

import (
	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sLog"
)

func getPriceStep(priceFrom, priceTo, stepPercentage float64) (float64, int) {

	width := sConv.GetDiffPercentAbs(priceFrom, priceTo)

	widthInt := int64(width * 100)
	stepPercentageInt := int64(stepPercentage * 100)

	quant := widthInt / stepPercentageInt

	if (widthInt % stepPercentageInt) > 0 {
		quant++
	}

	return sConv.GetDiffAbs(priceFrom, priceTo) / float64(quant), int(quant)
}

func CalcOrderPrices(priceFrom, priceTo, stepPercentage, profitPercentage, stopLossPercentage float64) (pricesIn, pricesOut []float64) {

	priceIn := priceFrom
	priceStep, quant := getPriceStep(priceFrom, priceTo, stepPercentage)

	if priceFrom > priceTo {
		for i := 0; i <= quant; i++ {
			pricesIn = append(pricesIn, priceIn)
			pricesOut = append(pricesOut, sConv.GetWithPercent(priceIn, profitPercentage, true))
			priceIn = priceIn - priceStep
		}
	} else {
		for i := 0; i <= quant; i++ {
			pricesIn = append(pricesIn, priceIn)
			pricesOut = append(pricesOut, sConv.GetWithPercent(priceIn, profitPercentage, false))
			priceIn = priceIn + priceStep
		}
	}
	return
}

func CalcOrderSizes(pricesIn []float64, balance float64) (sizes []float64) {

	count := len(pricesIn)
	size := balance / float64(count)

	for i := 0; i < count; i++ {
		sizes = append(sizes, size)
	}
	return
}

func CalcOrderSizesWeighted(pricesIn []float64, balance float64) (sizes []float64) {

	count := len(pricesIn)
	mean := sConv.GetDiffAbs(pricesIn[0], pricesIn[count-1])

	if pricesIn[0] > pricesIn[count-1] {
		mean = pricesIn[0] - (mean / 2)
	} else {
		mean = pricesIn[0] + (mean / 2)
	}
	sLog.Debug("mean: %.4f", mean)

	contracts := balance / mean / float64(count)
	for _, price := range pricesIn {
		sizes = append(sizes, (contracts * price))
	}
	return
}

func CalcOrderSizesStep(pricesIn []float64, balance, step float64) (sizes []float64) {

	count := len(pricesIn)
	div := float64(count / 2)
	first := 100.0 / float64(count)

	if (count % 2) == 0 {
		first += (div - 0.5) * step
	} else {
		first += div * step
	}

	for i := 0; i < count; i++ {
		sizes = append(sizes, (balance * first / 100))
		first -= step
	}
	return
}

func CalcStopLossOut(pricesIn, sizes []float64, stopLossPercentage, balance float64) float64 {

	count := len(pricesIn)
	if count == 0 || count != len(sizes) {
		return 0
	}

	var contracts float64
	for k, price := range pricesIn {
		contracts += sizes[k] / price
	}

	price := balance / contracts

	if pricesIn[0] > pricesIn[count-1] {
		return sConv.GetWithPercent(price, stopLossPercentage, false)
	}
	return sConv.GetWithPercent(price, stopLossPercentage, true)
}

func CalcProfitOut(pricesIn, sizes []float64, stopLossPercentage, balance float64) (pricesOut []float64) {

	count := len(pricesIn)
	if count == 0 || count != len(sizes) {
		return
	}

	var contracts float64
	for k, price := range pricesIn {
		contracts += sizes[k] / price
	}

	price := balance / contracts

	if pricesIn[0] > pricesIn[count-1] {
		sConv.GetWithPercent(price, stopLossPercentage, false)
	}
	sConv.GetWithPercent(price, stopLossPercentage, true)

	return
}

func CalcProfitAndLoss(pricesIn, pricesOut, sizes []float64, side int) (totalReturn float64, slice []float64) {

	switch side {
	case Sell:
		for i := range pricesIn {
			pnl := (pricesIn[i] - pricesOut[i]) * sizes[i] / pricesIn[i]
			slice = append(slice, pnl)
			totalReturn += pnl
		}
	case Buy:
		for i := range pricesIn {
			pnl := (pricesOut[i] - pricesIn[i]) * sizes[i] / pricesIn[i]
			slice = append(slice, pnl)
			totalReturn += pnl
		}
	}
	return
}

func CalcOrderPricesSizes(priceFrom, priceTo, stepPercentage, profitPercentage, stopLossPercentage, balance, sizeStepPercentage float64, weighted, unifiedProfit bool) (pricesIn, pricesOut, sizes []float64, stopLoss float64) {

	pricesIn, pricesOut = CalcOrderPrices(priceFrom, priceTo, stepPercentage, profitPercentage, stopLossPercentage)

	if sizeStepPercentage > 0 {
		sizes = CalcOrderSizesStep(pricesIn, balance, sizeStepPercentage)
	} else if weighted {
		sizes = CalcOrderSizesWeighted(pricesIn, balance)
	} else {
		sizes = CalcOrderSizes(pricesIn, balance)
	}

	stopLoss = CalcStopLossOut(pricesIn, sizes, stopLossPercentage, balance)

	if unifiedProfit {
		pricesOut = CalcProfitOut(pricesIn, sizes, stopLossPercentage, balance)
	}
	return
}
