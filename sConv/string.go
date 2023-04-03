package sConv

import (
	"strconv"
	"strings"
)

func GetBool(str string) bool {

	v, _ := strconv.ParseBool(str)

	return v
}

func GetStrB(v bool) string {
	return strconv.FormatBool(v)
}

func GetStrF(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}

func GetStrI(v int) string {
	return strconv.Itoa(v)
}

func GetStrI64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func GetStrPrecF(v float64, prec int) string {
	return strconv.FormatFloat(v, 'f', prec, 64)
}

func GetStrArrF(vs []float64) string {

	res := []string{}
	for _, v := range vs {
		res = append(res, strconv.FormatFloat(v, 'f', 2, 64))
	}

	return "[" + strings.Join(res, ",") + "]"
}
