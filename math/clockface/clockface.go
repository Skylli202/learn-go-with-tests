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

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func secondHandPoint(t time.Time) Point {
	a := secondsInRadians(t)
	return angleToPoint(a)
}

func minuteHandPoint(t time.Time) Point {
	a := minutesInRadians(t)
	return angleToPoint(a)
}

func angleToPoint(a float64) Point {
	x := math.Sin(a)
	y := math.Cos(a)

	return Point{x, y}
}
