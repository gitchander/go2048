package core

import "image"

type Point = image.Point

func MakePoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}
