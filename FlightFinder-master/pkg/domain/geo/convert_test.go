package geo_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/geo"
)

func TestConvertLongitude(t *testing.T) {
	// given
	cases := []struct {
		deg, min, sec int
		hem           string
		expected      geo.Longitude
		err           error
	}{
		{10, 0, 0, "E", 10.0, nil},
		{10, 30, 0, "E", 10.5, nil},
		{10, 30, 36, "E", 10.51, nil},
		{10, 0, 0, "W", -10.0, nil},
		{10, 30, 0, "W", -10.5, nil},
		{10, 30, 36, "W", -10.51, nil},
		{10, 30, 36, "P", 0.0, geo.ConversionError{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking %v", c), func(t *testing.T) {
			// when
			actual, err := geo.ConvertDegMinSecHemToLongitude(c.deg, c.min, c.sec, c.hem)
			if bool(c.err == nil) != bool(err == nil) {
				t.Errorf("Expected conversion error, got %v", err)
			}
			if c.expected != actual {
				t.Errorf("Expected longitude %v, got %v", c.expected, actual)
			}
		})
	}
}

func TestConvertLatitude(t *testing.T) {
	// given
	cases := []struct {
		deg, min, sec int
		hem           string
		expected      geo.Latitude
		err           error
	}{
		{10, 0, 0, "N", 10.0, nil},
		{10, 30, 0, "N", 10.5, nil},
		{10, 30, 36, "N", 10.51, nil},
		{10, 0, 0, "S", -10.0, nil},
		{10, 30, 0, "S", -10.5, nil},
		{10, 30, 36, "S", -10.51, nil},
		{10, 30, 36, "P", 0.0, geo.ConversionError{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking %v", c), func(t *testing.T) {
			// when
			actual, err := geo.ConvertDegMinSecHemToLatitude(c.deg, c.min, c.sec, c.hem)
			if bool(c.err == nil) != bool(err == nil) {
				t.Errorf("Expected conversion error, got %v", err)
			}
			if c.expected != actual {
				t.Errorf("Expected latitude %v, got %v", c.expected, actual)
			}
		})
	}
}
