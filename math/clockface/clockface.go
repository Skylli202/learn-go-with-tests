package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock

	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock

	hoursInHalfClock = 6
	hoursInClock     = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func secondHandPoint(t time.Time) Point {
	a := secondsInRadians(t)
	return angleToPoint(a)
}

func minuteHandPoint(t time.Time) Point {
	a := minutesInRadians(t)
	return angleToPoint(a)
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(a float64) Point {
	x := math.Sin(a)
	y := math.Cos(a)

	return Point{x, y}
}
