package postprocessing_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
)

func TestCSVnationsFilterReturnsValidNations(t *testing.T) {
	// given
	csvNations := make(chan loading.CSVNation, 3)
	csvNations <- loading.NewCSVNation("CN", "CHN", "CNY", "CHINA")
	csvNations <- loading.NewCSVNation("ES", "ESP", "EUR", "SPAIN")
	csvNations <- loading.NewCSVNation("PL", "POL", "PLN", "POLAND")
	close(csvNations)
	// notice: expectedNations are sorted ascending
	expectedNations := nations.Nations{
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	}

	// when
	actualNations := postprocessing.FilterNations(csvNations)

	// then
	if len(actualNations) != len(expectedNations) {
		t.Errorf("Expected num nations %d, got %d", len(expectedNations), len(actualNations))
	}

	for i, actual := range actualNations {
		if actual != expectedNations[i] {
			t.Errorf("At index %d expected nation %v, got %v", i, expectedNations[i], actual)
		}
	}
}
