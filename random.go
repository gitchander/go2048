package go2048

import (
	"math"
	"math/rand"
	"time"
)

var random = func() func() float64 {
	values := make(chan float64)
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			values <- r.Float64()
		}
	}()
	return func() float64 {
		return <-values
	}
}()

func randIntn(n int) int {
	return int(math.Floor(random() * float64(n)))
}
