package sConv

import (
	"log"
	"sort"
	"strconv"

	"github.com/yasseldg/simplego/unique"
)

type I struct{}

func (I) isType(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Supplied value %s is not a integer. Error: %s", str, err)
		return false
	}
	return true
}

func (I) getRange(extV ExtValues) []string {

	V_1, _ := strconv.Atoi(extV.V_1)
	V_2 := int(0)
	if len(extV.V_2) > 0 {
		V_2, _ = strconv.Atoi(extV.V_2)
	}
	step := int(0)
	if len(extV.step) > 0 {
		step, _ = strconv.Atoi(extV.step)
	}

	values := []int{}

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

	conV := []string{}
	for _, cV := range values {
		conV = append(conV, strconv.Itoa(cV))
	}

	return conV
}

func (I) validateRangeLimits(extV ExtValues) bool {

	V_1, _ := strconv.Atoi(extV.V_1)
	V_2 := int(0)
	if len(extV.V_2) > 0 {
		V_2, _ = strconv.Atoi(extV.V_2)
	}

	return (V_2 > V_1)
}

func (I) sortRange(values []string) []string {

	strVs := []string{}
	intVs := []int{}

	for _, v := range values {
		intV, _ := strconv.Atoi(v)
		intVs = append(intVs, intV)
	}

	sort.Ints(intVs)
	unique.Ints(&intVs)

	for _, v := range intVs {
		strVs = append(strVs, strconv.Itoa(v))
	}

	return strVs
}

func (I) compare(op, v1, v2 string) bool {

	cv1 := GetInt(v1)
	cv2 := GetInt(v2)

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

func GetInt(str string) int {

	v, _ := strconv.Atoi(str)

	return v
}

func GetInt64(str string) int64 {

	v, _ := strconv.ParseInt(str, 10, 64)

	return v
}
