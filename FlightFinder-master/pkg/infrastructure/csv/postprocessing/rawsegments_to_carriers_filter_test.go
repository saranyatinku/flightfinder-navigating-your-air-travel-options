package postprocessing_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"
)

func TestAirportsCarriersLoaderReturnsValidData2(t *testing.T) {
	// given
	csvSegments := make(chan loading.CSVSegment, 3)
	csvSegments <- loading.NewCSVSegment("KRK", "WAW", "LO")
	csvSegments <- loading.NewCSVSegment("WAW", "WRO", "LH")
	csvSegments <- loading.NewCSVSegment("WRO", "GDN", "BY")
	close(csvSegments)
	// expected carriers are sorted
	expectedCarrires := carriers.Carriers{
		carriers.NewCarrier("BY"),
		carriers.NewCarrier("LH"),
		carriers.NewCarrier("LO"),
	}

	// when
	actualCarriers := postprocessing.ExtractCarriers(csvSegments)

	// then
	if len(expectedCarrires) != len(actualCarriers) {
		t.Errorf("Expected num carriers %d, got %d", len(expectedCarrires), len(actualCarriers))
	}

	for i, actual := range actualCarriers {
		if actual != actualCarriers[i] {
			t.Errorf("At index %d expected carrier %v, got %v", i, expectedCarrires[i], actual)
		}
	}
}
