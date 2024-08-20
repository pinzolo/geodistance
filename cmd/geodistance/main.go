package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/pinzolo/geodistance"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const (
	modeGeographicLib = "geographiclib"
	modeHaversine     = "haversine"
	modeFlat          = "flat"
)

var (
	mode  string
	debug bool
)

func main() {
	flag.StringVar(&mode, "mode", modeGeographicLib, "calculation mode (geographiclib, haversine, flat)")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	c, err := getCalculator()
	if err != nil {
		handleError(err)
	}
	args := flag.Args()
	if len(args) < 2 {
		handleError(errors.New("need two points"))
	}
	p1, err := parsePoint(args[0])
	if err != nil {
		handleError(err)
	}
	p2, err := parsePoint(args[1])
	if err != nil {
		handleError(err)
	}
	r := c.Calculate(p1, p2)
	fmt.Printf("%.5f\n", r)
}

func getCalculator() (geodistance.DistanceCalculator, error) {
	logger := newLogger()
	if mode == modeGeographicLib {
		return geodistance.NewGeographicLibCalculator(logger), nil
	} else if mode == modeHaversine {
		return geodistance.NewHaversineCalculator(logger), nil
	} else if mode == modeFlat {
		return geodistance.NewFlatCalculator(logger), nil
	}

	return nil, fmt.Errorf("invalid mode: %s", mode)
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

func parsePoint(s string) (geodistance.Point, error) {
	split := strings.Split(s, ",")
	if len(split) != 2 {
		return geodistance.Point{}, fmt.Errorf("invalid point: %s", s)
	}
	lat, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return geodistance.Point{}, err
	}
	lng, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return geodistance.Point{}, err
	}
	return geodistance.NewPoint(lat, lng)
}
