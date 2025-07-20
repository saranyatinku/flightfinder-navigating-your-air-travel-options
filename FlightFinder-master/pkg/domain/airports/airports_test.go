package airports_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
)

func TestNewAirportsReturnsAllAirportsSortedByCode(t *testing.T) {
	// given
	// expectedAirports are sorted ascending by code
	expectedAirports := airports.Airports{
		airports.NewAirportCodeOnly("AAA"),
		airports.NewAirportCodeOnly("KKK"),
		airports.NewAirportCodeOnly("ZZZ"),
	}

	// when
	actualAirports := airports.NewAirports(
		airports.NewAirportCodeOnly("KKK"),
		airports.NewAirportCodeOnly("AAA"),
		airports.NewAirportCodeOnly("ZZZ"),
	)

	// then
	if len(expectedAirports) != len(actualAirports) {
		t.Fatalf("Num expected airports != num actual airports: %d : %d", len(expectedAirports), len(actualAirports))
	}

	for i, expected := range expectedAirports {
		actual := actualAirports[i]
		if actual != expected {
			t.Errorf("At index %d expected airport %v, got %v", i, expected, actual)
		}
	}
}

func TestGetByCodeReturnsValidAirport(t *testing.T) {
	// given
	// important: airports are sorted ascending for binary search
	airportList := airports.NewAirports(
		airports.NewAirportCodeOnly("AAA"),
		airports.NewAirportCodeOnly("KKK"),
		airports.NewAirportCodeOnly("ZZZ"),
	)
	cases := []struct {
		code string
		id   airports.ID
	}{
		{"AAA", 0},
		{"GGG", airports.NullID},
		{"KKK", 1},
		{"PPP", airports.NullID},
		{"ZZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking ID for %s", c.code), func(t *testing.T) {
			// when
			id := airportList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected ID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
