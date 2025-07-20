package postprocessing

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

func ExtractAirports(segments <-chan loading.CSVSegment) airports.Airports {
	uniqueCodes := make(map[string]bool)
	for s := range segments {
		uniqueCodes[s.FromAirportCode] = true
		uniqueCodes[s.ToAirportCode] = true
	}

	var list []airports.Airport
	for code := range uniqueCodes {
		list = append(list, airports.NewAirportCodeOnly(code))
	}

	return airports.NewAirports(list...)
}
