package geodistance_test

import (
	"github.com/pinzolo/geodistance"
	"log/slog"
	"math/rand"
	"os"
	"testing"
)

func BenchmarkFlatCalculator(b *testing.B) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	c := geodistance.NewFlatCalculator(logger)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p1 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		p2 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		c.Calculate(p1, p2)
	}
}

func BenchmarkHaversineCalculator(b *testing.B) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	c := geodistance.NewHaversineCalculator(logger)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p1 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		p2 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		c.Calculate(p1, p2)
	}
}

func BenchmarkGeographicLibCalculator(b *testing.B) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	c := geodistance.NewGeographicLibCalculator(logger)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p1 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		p2 := geodistance.Point{
			Longitude: getRandomLongitude(),
			Latitude:  getRandomLatitude(),
		}
		c.Calculate(p1, p2)
	}
}

func getRandomLatitude() float64 {
	if rand.Intn(2) == 0 {
		return rand.Float64() * 90.0
	} else {
		return rand.Float64() * -90.0
	}
}

func getRandomLongitude() float64 {
	if rand.Intn(2) == 0 {
		return rand.Float64() * 180.0
	} else {
		return rand.Float64() * -180.0
	}
}
