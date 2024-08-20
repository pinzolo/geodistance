package geodistance

import (
	"github.com/paulmach/orb/geo"
	"log/slog"
)

type flatCalculator struct {
	logger *slog.Logger
}

func NewFlatCalculator(logger *slog.Logger) DistanceCalculator {
	return &flatCalculator{logger: logger}
}

func (c *flatCalculator) Calculate(p1 Point, p2 Point) float64 {
	c.logger.Debug("Method: flat", slog.String("point1", p1.String()), slog.String("point2", p2.String()))
	return geo.Distance(p1.Orb(), p2.Orb())
}
