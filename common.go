package geodistance

import (
	"fmt"
	"github.com/paulmach/orb"
	"strconv"
	"strings"
)

const (
	R                   = 6378137.0
	MethodFlat          = "flat"
	MethodHaversine     = "haversine"
	MethodGeographicLib = "geographiclib"
)

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
	Method() string
	Calculate(p1 Point, p2 Point) float64
}

func ParsePoint(s string) (Point, error) {
	split := strings.Split(s, ",")
	if len(split) != 2 {
		return Point{}, fmt.Errorf("invalid point: %s", s)
	}
	lat, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return Point{}, err
	}
	lng, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return Point{}, err
	}
	return NewPoint(lat, lng)
}
