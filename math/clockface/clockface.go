package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	a := secondsInRadians(t)

	x := math.Sin(a)
	y := math.Cos(a)

	return Point{x, y}
}
