package sSlice

func InStrings(needle string, values []string) bool {
	return InStringsPos(needle, values) >= 0
}

func InStringsPos(needle string, values []string) int {
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
