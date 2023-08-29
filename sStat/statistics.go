package sStat

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/stat/distuv"

	xrand "golang.org/x/exp/rand"
)

// PoissonDist
func PoissonDist(lambda, min, max float64, n int) (dist []float64) {

	if min <= lambda && lambda <= max {
		source := xrand.NewSource(uint64(time.Now().UnixNano()))

		poisson := distuv.Poisson{
			Lambda: lambda,
			Src:    source,
		}

		for n > 0 {
			r := poisson.Rand()
			if min <= r && r <= max {
				dist = append(dist, r)
				n--
			}
		}
	}
	return
}

// UniformDist, base <= (rand float64) < 1
func UniformDist(min float64) float64 {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10000; i++ {
		r := rand.Float64()
		if min <= r && r < 1.0 {
			return r
		}
	}
	return min
}

func ValidFloat64(f float64) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0
	}
	return f
}

func Zscore(value, mean, std_dev float64) float64 {
	return ValidFloat64((math.Abs(value) - math.Abs(mean)) / math.Abs(std_dev))
}

func ZscoreAbs(value, mean, std_dev float64) float64 {
	return math.Abs(Zscore(value, mean, std_dev))
}

// Perc Category
func PercCategory(value, min, max float64, n int) int {
	if value <= min {
		return 1
	}
	if value >= max {
		return n
	}
	return int((value-min)/PercStep(min, max, n)) + 1
}

func PercStep(min, max float64, n int) float64 {
	return (max - min) / float64(n)
}
