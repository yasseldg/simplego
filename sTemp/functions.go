package sTemp

import (
	"fmt"
	"text/template"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sDate"
)

func Functions(name ...string) template.FuncMap {
	return template.FuncMap{
		"Float": func(f float64, prec int) string {
			return FormatFloat(f, prec)
		},

		"FormatF": func(f float64, dec string) string {
			return FormatF(f, dec)
		},

		"MultF": func(v1, v2 float64, dec string) string {
			prec := "%." + dec + "f"
			return fmt.Sprintf(prec, (v1 * v2))
		},

		"MultI": func(v1, v2 int64) string {

			return fmt.Sprintf("%d", (v1 * v2))
		},

		"RestF": func(v1, v2 int) int {
			return v1 - v2
		},

		"RestI": func(v1, v2 int) int {
			return v1 % v2
		},

		"DivI": func(v1, v2 int) int {
			return v1 / v2
		},

		"SumI": func(v1, v2 int) int {
			return v1 + v2
		},

		"SumI64": func(v1, v2 int64) int64 {
			return v1 + v2
		},

		"GetI": func(s string) int {
			return sConv.GetInt(s)
		},

		"FormatD": func(value any, prec int64) string {
			return sDate.ForWeb(value, prec)
		},

		"FormatDform": func(value any, prec int64) string {
			return sDate.FormatDSep(value, prec, "-")
		},

		"GetValues": func(str, vt string) []string {
			return sConv.GetValues(str, vt)
		},

		"GetRangeLimits": func(str, vt string) []sConv.ExtValues {
			return sConv.GetRangeLimits(str, vt)
		}}
}

func FormatFloat(f float64, prec int) string {
	if prec < 0 {
		return fmt.Sprintf("%f", f)
	}
	return fmt.Sprintf(fmt.Sprintf("%%.%df", prec), f)
}

//  ----  OLD verion ----

func FormatF(f float64, dec string) string {
	prec := "%." + dec + "f"
	return fmt.Sprintf(prec, f)
}
