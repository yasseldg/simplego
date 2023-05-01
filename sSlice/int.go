package sSlice

func SumI(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return
}

func CountI(values []int) map[int]int {
	counts := make(map[int]int)
	for _, v := range values {
		counts[v]++
	}
	return counts
}

func InInts(needle int, values []int) bool {
	return InIntsPos(needle, values) >= 0
}

func InIntsPos(needle int, values []int) int {
	for i, v := range values {
		if v == needle {
			return i
		}
	}
	return -1
}
