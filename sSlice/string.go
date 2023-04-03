package sSlice

func InStrings(needle string, values []string) bool {
	for _, v := range values {
		if v == needle {
			return true
		}
	}
	return false
}
