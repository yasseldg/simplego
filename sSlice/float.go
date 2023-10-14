package sSlice

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/unique"
)

// SortUniqueMinMax and delete values "0"
func SortUniqueMinMaxF(values []float64) (float64, float64) {

	sort.Float64s(values)
	unique.Float64s(&values)

	min := 0.0
	max := 0.0
	tValues := []float64{}

	if len(values) > 0 {
		min = values[0]
		if min == 0 {
			tValues = append(tValues, values[1:]...)
		}

		if len(values) > 0.0 {
			min = values[0]
			max = values[len(values)-1]
		}
	}

	values = tValues

	return min, max
}

// Min and Max positive values of the float slice, including "0" or not
func MinMaxF(values []float64, zero bool) (float64, float64) {

	min := math.Inf(1)
	max := math.Inf(-1)

	for _, v := range values {
		if v > max {
			max = v
		}

		if v < min {
			if !zero && v == 0.0 {
				continue
			}
			min = v
		}
	}

	return min, max
}

func CompareFloat(s1, s2 []float64, prec int, diff float64) bool {
	out_set := len(s2) - len(s1)
	for k, v1 := range s1 {
		sub := math.Abs(sConv.CompareTruncFloat(v1, s2[k+out_set], prec))
		if sub > 0 && sub > diff {
			return false
		}
	}
	return true
}

func CompareReverseFloat(s1, s2 []float64, prec int, diff float64) bool {
	l := len(s2)
	out_set := l - len(s1)
	for k, v1 := range s1 {
		sub := math.Abs(sConv.CompareTruncFloat(v1, s2[l-1-k-out_set], prec))
		if sub > 0 && sub > diff {
			return false
		}
	}
	return true
}

// SortF, order "0=asc","1=desc"
func SortF(values []float64, order int) {
	switch order {
	case 0:
		sort.Sort(sort.Float64Slice(values))
	case 1:
		// ordenamos los precios de mayor a menor
		sort.Sort(sort.Reverse(sort.Float64Slice(values)))
	}
}

// SortFloat64, order "1=asc","-1=desc"
func SortFloat64(values []float64, order int) {
	switch order {
	case 1:
		sort.Sort(sort.Float64Slice(values))
	case -1:
		sort.Sort(sort.Reverse(sort.Float64Slice(values)))
	}
}

// TrimF, trim greater or lesser, order "0=asc: greater","1=desc: lesser"
func TrimF(values []float64, keepOut float64, order int) []float64 {
	// log.Printf("TrimSliceFloat: keepOut: %.2f  --  values: %v", keepOut, values)
	for k, price := range values {
		switch order {
		case 0:
			if price > keepOut {
				// log.Printf("TrimSliceFloat: keepOut: %.2f  --  values[k:]: %v", keepOut, values[k:])
				return values[k:]
			}
		case 1:
			if price < keepOut {
				// log.Printf("TrimSliceFloat: keepOut: %.2f  --  values[k:]: %v", keepOut, values[k:])
				return values[k:]
			}
		}
	}
	return []float64{}
}

// TrimLastF, trim greater or lesser, order "0=asc: greater","1=desc: lesser"
func TrimLastF(values []float64, keepOut float64, order int) []float64 {
	// log.Printf("TrimSliceFloat: keepOut: %.2f  --  values: %v", keepOut, values)
	for k := len(values) - 1; k >= 0; k-- {
		switch order {
		case 0:
			if values[k] < keepOut {
				// log.Printf("TrimSliceFloat: values[:k+1]: %v", values[:k+1])
				return values[:k+1]
			}
		case 1:
			if values[k] > keepOut {
				// log.Printf("TrimSliceFloat: values[:k+1]: %v", values[:k+1])
				return values[:k+1]
			}
		}
	}
	return []float64{}
}

func SumF(slice []float64) (sum float64) {
	for _, v := range slice {
		sum += v
	}
	return
}

func CountF(values []float64) map[float64]int {

	counts := make(map[float64]int)

	for _, v := range values {
		counts[v]++
	}

	return counts
}

func WeightF(values []float64, prec int) map[float64]float64 {

	weights := make(map[float64]float64)

	counts := CountF(values)

	l := float64(len(values))
	for v, c := range counts {
		weights[v] = sConv.GetPrecFloat((float64(c) / l), prec)
	}

	return weights
}

func FloatsToString(s []float64, prec int) string {
	strs := make([]string, len(s))
	for i, v := range s {
		strs[i] = fmt.Sprintf("%.*f", prec, v)
	}
	return strings.Join(strs, ", ")
}
