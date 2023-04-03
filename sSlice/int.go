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
