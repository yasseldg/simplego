package sSlice

func InStrings(needle string, values []string) bool {
	return InStringsPos(needle, values) >= 0
}

func InStringsPos(needle string, values []string) int {
	if len(values) == 0 {
		return -1
	}

	for i, v := range values {
		if v == needle {
			return i
		}
	}
	return -1
}

func MapNames(values map[string]any) []string {
	names := []string{}
	for v := range values {
		names = append(names, v)
	}
	return names
}

func DeleteStrings(values []string, needle string) []string {
	if len(values) == 0 {
		return values
	}

	pos := InStringsPos(needle, values)
	if pos < 0 {
		return values
	}

	if len(values) == 1 {
		return []string{}
	}

	if pos == len(values)-1 {
		return values[:pos]
	}

	return append(values[:pos], values[pos+1:]...)
}
