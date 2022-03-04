// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/math

package clockface_svg

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30) // By dividing math.Pi by 30 and then by multiplying it by 30 we've ended up with a number that's no longer the same as math.Pi.
	return (math.Pi / (30 / (float64(t.Second()))))
}
