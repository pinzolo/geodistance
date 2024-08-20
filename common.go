package geodistance

import (
	"fmt"
	"github.com/paulmach/orb"
	"math"
)

const (
	R = 6378137.0
)

func Deg2Rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

type Point struct {
	Latitude  float64
	Longitude float64
}

func (p Point) String() string {
	return fmt.Sprintf("{%v, %v}", p.Latitude, p.Longitude)
}

func (p Point) Orb() orb.Point {
	return orb.Point{p.Latitude, p.Longitude}
}

func NewPoint(lat, lng float64) (Point, error) {
	if lat < -90 || 90 < lat {
		return Point{}, fmt.Errorf("latitude must be between -90 and 90 degrees: %v", lat)
	}
	if lng < -180 || 180 < lng {
		return Point{}, fmt.Errorf("longitude must be between -180 and 180 degrees: %v", lng)
	}

	return Point{lat, lng}, nil
}

type DistanceCalculator interface {
	Calculate(p1 Point, p2 Point) float64
}
