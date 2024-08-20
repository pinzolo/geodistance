package geodistance

import (
	"github.com/paulmach/orb/geo"
	"log/slog"
)

type haversineCalculator struct {
	logger *slog.Logger
}

func NewHaversineCalculator(logger *slog.Logger) DistanceCalculator {
	return &haversineCalculator{logger: logger}
}

func (c *haversineCalculator) Calculate(p1 Point, p2 Point) float64 {
	c.logger.Debug("Method: haversine", slog.String("point1", p1.String()), slog.String("point2", p2.String()))
	return geo.DistanceHaversine(p1.Orb(), p2.Orb())
}
