package sConv

import (
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/yasseldg/simplego/unique"
)

type F struct{}

func (F) isType(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Printf("isType. Supplied value %s is not a float. Error: %s", str, err)
		return false
	}
	return true
}

func (F) getRange(extV ExtValues) []string {

	V_1, _ := strconv.ParseFloat(extV.V_1, 64)
	V_2 := float64(0)
	if len(extV.V_2) > 0 {
		V_2, _ = strconv.ParseFloat(extV.V_2, 64)
	}
	step := float64(0)
	if len(extV.step) > 0 {
		step, _ = strconv.ParseFloat(extV.step, 64)
	}

	values := []float64{}

	if V_2 > V_1 {
		if step >= 0 {
			if step == 0 {
				step = 1
			}

			nV := V_1
			for nV <= V_2 {
				values = append(values, nV)
				nV += step
			}
		}
	} else {
		values = append(values, V_1)
	}

	sort.Float64s(values)
	// sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

	conV := []string{}
	for _, cV := range values {
		conV = append(conV, strconv.FormatFloat(cV, 'f', 2, 64))
	}

	return conV
}

func (F) validateRangeLimits(extV ExtValues) bool {

	V_1, _ := strconv.ParseFloat(extV.V_1, 64)
	V_2 := float64(0)
	if len(extV.V_2) > 0 {
		V_2, _ = strconv.ParseFloat(extV.V_2, 64)
	}

	return (V_2 > V_1)
}

func (F) sortRange(values []string) []string {

	floatVs := []float64{}
	for _, v := range values {
		floatVs = append(floatVs, GetFloat(v))
	}

	sort.Float64s(floatVs)
	unique.Float64s(&floatVs)

	strVs := []string{}
	for _, v := range floatVs {
		strVs = append(strVs, GetStrF(v))
	}

	return strVs
}

func (F) compare(op, v1, v2 string) bool {

	cv1 := GetFloat(v1)
	cv2 := GetFloat(v2)

	switch op {
	case ">":
		return cv1 > cv2
	case ">=":
		return cv1 >= cv2
	case "<":
		return cv1 < cv2
	case "<=":
		return cv1 <= cv2
	case "==":
		return cv1 == cv2
	}

	return false
}

func (F) RemoveIndex(slice []float64, index int) []float64 {

	if len(slice) > index && index >= 0 {
		return append(slice[:index], slice[index+1:]...)
	}

	return slice
}

func (F) RemoveValue(slice []float64, value float64) []float64 {

	fV := F{}
	for k, v := range slice {
		if v == value {
			return fV.RemoveIndex(slice, k)
		}
	}

	return slice
}

// func (F) Average(slice []float64) float64 {

// 	return sSlice.SumSliceFloat(slice) / float64(len(slice))
// }

func GetFloat(str string) float64 {

	v, _ := strconv.ParseFloat(str, 64)

	return v
}

func GetPrecFloat(value float64, prec int) float64 {
	pow := math.Pow10(prec)
	return math.Round(value*pow) / pow
}

func GetTruncFloat(value float64, prec int) float64 {
	pow := math.Pow10(prec)
	return float64(int64(value*pow)) / pow
}

// ComparePrecFloat, "v1>v2: >0",  "v1<v2: <0",  "v1=v2: 0"
func ComparePrecFloat(v1, v2 float64, prec int) float64 {
	return GetPrecFloat(v1, prec) - GetPrecFloat(v2, prec)
}

// CompareTruncFloat, "v1>v2: >0",  "v1<v2: <0",  "v1=v2: 0"
func CompareTruncFloat(v1, v2 float64, prec int) float64 {
	return GetTruncFloat(v1, prec) - GetTruncFloat(v2, prec)
}

func GetDiffPercent(fromValue, toValue float64) float64 {

	return ((toValue / fromValue) - 1) * 100
}

func GetDiffPercentAbs(fromValue, toValue float64) float64 {

	return math.Abs(GetDiffPercent(fromValue, toValue))
}

func GetDiffAbs(fromValue, toValue float64) float64 {

	return math.Abs(fromValue - toValue)
}

// Adding or Subtracting percent to value
func GetWithPercent(value, percent float64, adding bool) float64 {

	if adding {
		return ((percent / 100) + 1) * value
	}
	return -((percent / 100) - 1) * value
}

func GetFloatInterface(i interface{}) (f float64) {
	switch i := i.(type) {
	case float64:
		f = i
	case int:
		f = float64(i)
	case string:
		f = GetFloat(i)
	default:
		f = 0.0
	}
	return
}
