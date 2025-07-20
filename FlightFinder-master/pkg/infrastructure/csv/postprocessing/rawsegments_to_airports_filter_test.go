package postprocessing_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
)

func TestAirportsFilterReturnsValidAirports(t *testing.T) {
	// given
	csvSegments := make(chan loading.CSVSegment, 3)
	csvSegments <- loading.NewCSVSegment("KRK", "WAW", "LO")
	csvSegments <- loading.NewCSVSegment("WAW", "WRO", "LH")
	csvSegments <- loading.NewCSVSegment("WRO", "GDN", "BY")
	close(csvSegments)
	// expected airports are sorted
	expectedAirports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"),
		airports.NewAirportCodeOnly("KRK"),
		airports.NewAirportCodeOnly("WAW"),
		airports.NewAirportCodeOnly("WRO"),
	}

	// when
	actualAirports := postprocessing.ExtractAirports(csvSegments)

	// then
	if len(expectedAirports) != len(actualAirports) {
		t.Errorf("Expected num airports %d, got %d", len(expectedAirports), len(actualAirports))
	}

	for i, actual := range actualAirports {
		if actual != expectedAirports[i] {
			t.Errorf("At index %d expected airport %v, got %v", i, expectedAirports[i], actual)
		}
	}
}
