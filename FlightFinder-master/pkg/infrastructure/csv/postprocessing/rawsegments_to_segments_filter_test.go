package postprocessing_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"
)

func TestCSVsegmentsFilterReturnsValidSegments(t *testing.T) {
	// given
	// important: airports are sorted
	airports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"), // ID=0
		airports.NewAirportCodeOnly("KRK"), // ID=1
		airports.NewAirportCodeOnly("WAW"), // ID=2
		airports.NewAirportCodeOnly("WRO"), // ID=3
	}
	// important: carrierrs are sorted
	carriers := carriers.Carriers{
		carriers.NewCarrier("BY"), // carrierID=0
		carriers.NewCarrier("LH"), // carrierID=1
		carriers.NewCarrier("LO"), // carrierID=2
	}
	// important: expected segments are sorted
	expectedSegments := segments.Segments{
		segments.NewSegment(1, 2, 2),
		segments.NewSegment(2, 3, 1),
		segments.NewSegment(3, 0, 0),
	}
	csvSegments := make(chan loading.CSVSegment, 3)
	csvSegments <- loading.NewCSVSegment("KRK", "WAW", "LO")
	csvSegments <- loading.NewCSVSegment("WAW", "WRO", "LH")
	csvSegments <- loading.NewCSVSegment("WRO", "GDN", "BY")
	close(csvSegments)
	filter := postprocessing.NewCSVSegmentsToSegmentsFilter(airports, carriers)

	// when
	actualSegments := filter.Filter(csvSegments)

	// then
	if len(actualSegments) != len(expectedSegments) {
		t.Errorf("Expected num segments %d, got %d", len(expectedSegments), len(actualSegments))
	}

	for i, actual := range actualSegments {
		if actual != expectedSegments[i] {
			t.Errorf("At index %d expected segment %v, got %v", i, expectedSegments[i], actual)
		}
	}
}
