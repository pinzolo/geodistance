package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/pinzolo/geodistance"
	"log/slog"
	"math"
	"os"
)

var (
	debug  bool
	origin string
)

type Result struct {
	Method   string  `json:"method"`
	Distance float64 `json:"distance"`
	Diff     float64 `json:"diff"`
	DiffRate float64 `json:"diff_rate"`
}

func main() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		handleError(errors.New("need two points"))
	}
	p1, err := geodistance.ParsePoint(args[0])
	if err != nil {
		handleError(err)
	}
	p2, err := geodistance.ParsePoint(args[1])
	if err != nil {
		handleError(err)
	}

	logger := newLogger()
	base := geodistance.NewGeographicLibCalculator(logger)
	calcs := []geodistance.DistanceCalculator{
		geodistance.NewHaversineCalculator(logger),
		geodistance.NewFlatCalculator(logger),
	}

	baseResult := Result{
		Method:   base.Method(),
		Distance: base.Calculate(p1, p2),
		Diff:     0,
		DiffRate: 0,
	}

	results := make([]Result, len(calcs))
	for i, c := range calcs {
		distance := c.Calculate(p1, p2)
		diff := math.Abs(baseResult.Distance - distance)
		results[i] = Result{
			Method:   c.Method(),
			Distance: distance,
			Diff:     diff,
			DiffRate: diff * 100 / distance,
		}
	}

	fmt.Println("| method | distance | diff | diff_rate |")
	fmt.Println("| --- | --- | --- | --- |")
	fmt.Printf("| %s | %v | %v | %v |\n", baseResult.Method, baseResult.Distance, baseResult.Diff, baseResult.DiffRate)
	for _, r := range results {
		fmt.Printf("| %s | %v | %v | %v |\n", r.Method, r.Distance, r.Diff, r.DiffRate)
	}
}

func newLogger() *slog.Logger {
	var opts *slog.HandlerOptions
	if debug {
		opts = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
