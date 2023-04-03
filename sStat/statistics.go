package sStat

import (
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
