package postprocessing_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/geo"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"
)

func TestEnricherAddsExpectedInformationToAirports(t *testing.T) {
	// given
	csvAirports := make(chan loading.CSVAirport, 3)
	csvAirports <- loading.NewCSVAirport("WAW", "Warszawa", "PL", geo.Longitude(20.0), geo.Latitude(51.0))
	csvAirports <- loading.NewCSVAirport("GDN", "Gdańsk", "PL", geo.Longitude(21.0), geo.Latitude(52.0))
	csvAirports <- loading.NewCSVAirport("KRK", "Kraków", "PL", geo.Longitude(19.0), geo.Latitude(50.0))
	close(csvAirports)
	// airports need be sorted by code for binary search purpose
	actualAirports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"),
		airports.NewAirportCodeOnly("KRK"),
		airports.NewAirportCodeOnly("WAW"),
	}
	expectedAirports := airports.Airports{
		airports.NewAirport("GDN", "Gdańsk", "PL", geo.Longitude(21.0), geo.Latitude(52.0)),
		airports.NewAirport("KRK", "Kraków", "PL", geo.Longitude(19.0), geo.Latitude(50.0)),
		airports.NewAirport("WAW", "Warszawa", "PL", geo.Longitude(20.0), geo.Latitude(51.0)),
	}

	// when
	postprocessing.EnrichAirports(actualAirports, csvAirports)

	// then
	for i, actual := range actualAirports {
		if actual != expectedAirports[i] {
			t.Errorf("For index %d expected airport %+v, got %+v", i, expectedAirports[i], actual)
		}
	}
}
