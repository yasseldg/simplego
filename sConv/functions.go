package sConv

import (
	"strings"
)

func isType(str, vType string) bool {

	switch vType {
	case "float":
		fV := F{}
		return fV.isType(str)

	case "int":
		iV := I{}
		return iV.isType(str)
	}

	return false
}

// Compare is "true" if (v1 "op" v2)
func Compare(op, v1, v2, vType string) bool {

	switch vType {
	case "float":
		fV := F{}
		return fV.compare(op, v1, v2)

	case "int":
		iV := I{}
		return iV.compare(op, v1, v2)
	}

	return false
}

func GetValues(str, vType string) []string {
	values := []string{}

	eValues := extractValues(str)

	for _, val := range eValues {

		// log.Printf("V_1: %s ... V_2: %s ... step: %s", val.V_1, val.V_2, val.step)

		if isType(val.V_1, vType) && (len(val.V_2) == 0 || isType(val.V_2, vType)) && (len(val.step) == 0 || isType(val.step, vType)) {

			valR := getRange(val, vType)

			values = append(values, valR...)

			// log.Printf("valR: %v", valR)
		}
	}

	values = sortRange(values, vType)

	return values
}

func GetRangeLimits(str, vType string) []ExtValues {

	values := []ExtValues{}

	eValues := extractValues(str)

	for _, val := range eValues {

		// log.Printf("V_1: %s ... V_2: %s ... step: %s", val.V_1, val.V_2, val.step)

		if isType(val.V_1, vType) && isType(val.V_2, vType) && validateRangeLimits(val, vType) {

			values = append(values, val)
		}
	}

	return values
}

func getRange(extV ExtValues, vType string) []string {

	switch vType {
	case "float":
		fV := F{}
		return fV.getRange(extV)

	case "int":
		iV := I{}
		return iV.getRange(extV)
	}

	return []string{}
}

func validateRangeLimits(extV ExtValues, vType string) bool {

	switch vType {
	case "float":
		fV := F{}
		return fV.validateRangeLimits(extV)

	case "int":
		iV := I{}
		return iV.validateRangeLimits(extV)
	}

	return false
}

func sortRange(values []string, vType string) []string {

	switch vType {
	case "float":
		fV := F{}
		return fV.sortRange(values)

	case "int":
		iV := I{}
		return iV.sortRange(values)
	}

	return []string{}
}

func extractValues(str string) []ExtValues {

	eValues := []ExtValues{}
	arrStr := strings.Split(str, ",")

	for _, value := range arrStr {

		value = strings.TrimSpace(value)

		extV := ExtValues{value, "", ""}

		if strings.Contains(value, "-") {

			arrValue := strings.Split(value, "-")
			extV.V_1 = strings.TrimSpace(arrValue[0])

			if strings.Contains(arrValue[1], ":") {
				arrV_2 := strings.Split(arrValue[1], ":")
				extV.V_2 = strings.TrimSpace(arrV_2[0])
				extV.step = strings.TrimSpace(arrV_2[1])
			} else {
				extV.V_2 = strings.TrimSpace(arrValue[1])
			}
		}

		eValues = append(eValues, extV)
	}

	return eValues
}
