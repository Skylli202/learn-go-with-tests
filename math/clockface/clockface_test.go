package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			want := c.angle

			if got != want {
				t.Errorf("Wanted %v radians, but got %v", want, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 0), Point{0, 1}},
	}

	for _, c := range cases {
		got := secondHandPoint(c.time)

		if !roughlyEqualPoint(got, c.point) {
			t.Errorf("Wanted %v Point, but got %v", c.point, got)
		}
	}
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func simpleTime(hour int, min int, sec int) time.Time {
	return time.Date(1234, time.January, 1, hour, min, sec, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
