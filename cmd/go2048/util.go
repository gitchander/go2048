package main

import (
	"math"
	"time"
)

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func round(x float64) int {
	return int(x + math.Copysign(0.5, x))
}

func lerp(v0, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

func lerpInt(v0, v1 int, t float64) int {
	if v0 == v1 {
		return v0
	}
	return round(lerp(float64(v0), float64(v1), t))
}

var getTime = func() func() float64 {
	begin := time.Now()
	return func() float64 {
		return time.Since(begin).Seconds()
	}
}()
