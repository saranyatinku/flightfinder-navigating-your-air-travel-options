package nations_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
)

func TestBuilderReturnsAllNationsSorted(t *testing.T) {
	// given
	// expectedNations are sorted ascending by code
	expectedNations := nations.Nations{
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	}

	// when
	actualNations := nations.NewNations(
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
	)

	// then
	if len(expectedNations) != len(actualNations) {
		t.Fatalf("Num expected nations != num actual nations: %d : %d", len(expectedNations), len(actualNations))
	}

	for i, expected := range expectedNations {
		actual := actualNations[i]
		if actual != expected {
			t.Errorf("At index %d expected nation was %v, got %v", i, expected, actual)
		}
	}
}

func TestGetByCodeReturnsValidNation(t *testing.T) {
	// given
	nationList := nations.NewNations(
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	)
	cases := []struct {
		code string
		id   nations.ID
	}{
		{"CN", 0},
		{"GG", nations.NullID},
		{"ES", 1},
		{"GG", nations.NullID},
		{"PL", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking ID for %s", c.code), func(t *testing.T) {
			// when
			id := nationList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected ID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
