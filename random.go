package go2048

import (
	"math/rand"
	"time"
)

type randomer interface {
	Intn(n int) int
	Float64() float64
}

func newRandNow() randomer {
	return newRandTime(time.Now())
}

func newRandTime(t time.Time) randomer {
	return newRandSeed(t.UTC().UnixNano())
}

func newRandSeed(seed int64) randomer {
	return rand.New(rand.NewSource(seed))
}
