package geodistance

import (
	"github.com/pymaxion/geographiclib-go/geodesic"
	"log/slog"
)

type geographicLibCalculator struct {
	logger *slog.Logger
}

func NewGeographicLibCalculator(logger *slog.Logger) DistanceCalculator {
	return &geographicLibCalculator{logger: logger}
}

func (c *geographicLibCalculator) Method() string {
	return MethodGeographicLib
}

func (c *geographicLibCalculator) Calculate(p1 Point, p2 Point) float64 {
	c.logger.Debug("Method: geographiclib", slog.String("point1", p1.String()), slog.String("point2", p2.String()))
	r := geodesic.WGS84.Inverse(p1.Latitude, p1.Longitude, p2.Latitude, p2.Longitude)
	return r.S12
}
