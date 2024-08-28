package geodistance

import (
	"log/slog"
	"math"
)

type haversineCalculator struct {
	logger *slog.Logger
}

func NewHaversineCalculator(logger *slog.Logger) DistanceCalculator {
	return &haversineCalculator{logger: logger}
}

func (c *haversineCalculator) Method() string {
	return MethodHaversine
}

func (c *haversineCalculator) Calculate(p1 Point, p2 Point) float64 {
	c.logger.Debug("Method: haversine", slog.String("point1", p1.String()), slog.String("point2", p2.String()))
	x1 := Deg2Rad(p1.Longitude)
	y1 := Deg2Rad(p1.Latitude)
	x2 := Deg2Rad(p2.Longitude)
	y2 := Deg2Rad(p2.Latitude)
	return R * math.Acos(math.Sin(y1)*math.Sin(y2)+math.Cos(y1)*math.Cos(y2)*math.Cos(x2-x1))
}
