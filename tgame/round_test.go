package main

import (
	"math"
	"testing"
)

type roundSample struct {
	param  float64
	result int
}

func TestRound(t *testing.T) {

	var samples = []roundSample{
		{-452.5, -453},
		{-452.4999999999999, -452},
		{-1.5, -2},
		{-1.499999999999999, -1},
		{-1, -1},
		{-0.5, -1},
		{-0.4999999999999999, 0},
		{-0.000000001, 0},
		{0, 0},
		{0.000000001, 0},
		{0.4999999999999999, 0},
		{0.5, 1},
		{1, 1},
		{1.499999999999999, 1},
		{1.5, 2},
		{723.4999999999999, 723},
		{723.5, 724},
	}

	testRoundFunc(t, "v1", roundV1, samples)
	testRoundFunc(t, "v2", roundV2, samples)
	testRoundFunc(t, "v3", roundV3, samples)
	testRoundFunc(t, "v4", roundV4, samples)
}

func testRoundFunc(t *testing.T, roundVersion string, round func(float64) int, samples []roundSample) {
	for _, sample := range samples {
		result := round(sample.param)
		if result != sample.result {
			t.Fatalf("round[%s](%g) = %d not equal sample result %d", roundVersion, sample.param, result, sample.result)
		}
	}
}

func roundV1(x float64) int {
	if x < 0 {
		return int(x - 0.5)
	}
	return int(x + 0.5)
}

func roundV2(x float64) int {
	if x <= -0.5 {
		return int(x - 0.5)
	}
	if x >= 0.5 {
		return int(x + 0.5)
	}
	return 0
}

func roundV3(x float64) int {
	if x < 0 {
		return int(math.Ceil(x - 0.5))
	}
	return int(math.Floor(x + 0.5))
}

func roundV4(x float64) int {
	return int(x + math.Copysign(0.5, x))
}
