package carriers_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
)

func TestNewCarriersReturnsAllCarriersSorted(t *testing.T) {
	// given
	// expectedCarriers are sorted ascending by code
	expectedCarriers := carriers.Carriers{
		carriers.NewCarrier("AA"),
		carriers.NewCarrier("BB"),
		carriers.NewCarrier("CC"),
	}

	// when
	actualCarriers := carriers.NewCarriers(
		carriers.NewCarrier("CC"),
		carriers.NewCarrier("BB"),
		carriers.NewCarrier("AA"),
	)

	// then
	if len(expectedCarriers) != len(actualCarriers) {
		t.Fatalf("Num expected airports != num actual airports: %d : %d", len(expectedCarriers), len(actualCarriers))
	}

	for i, expected := range expectedCarriers {
		actual := actualCarriers[i]
		if actual != expected {
			t.Errorf("At index %d expected carrier %v, got %v", i, expected, actual)
		}
	}
}

func TestGetByCodeReturnsValidCarrier(t *testing.T) {
	// given
	carrierList := carriers.NewCarriers(
		carriers.NewCarrier("AA"),
		carriers.NewCarrier("KK"),
		carriers.NewCarrier("ZZ"),
	)
	cases := []struct {
		code string
		id   carriers.ID
	}{
		{"AA", 0},
		{"GG", carriers.NullID},
		{"KK", 1},
		{"PP", carriers.NullID},
		{"ZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking CarrierID for %s", c.code), func(t *testing.T) {
			// when
			id := carrierList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected CarrierID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
