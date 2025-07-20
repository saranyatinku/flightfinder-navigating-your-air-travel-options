package segments_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
)

func TestBuilderReturnsAllSegmentsSorted(t *testing.T) {
	// given
	// airports must be sorted ascending
	airportList := airports.Airports{
		airports.NewAirportCodeOnly("AAA"), // ID 0
		airports.NewAirportCodeOnly("KKK"), // ID 1
		airports.NewAirportCodeOnly("ZZZ"), // ID 2
	}
	// carriers must be sorted ascending
	carriers := carriers.Carriers{
		carriers.NewCarrier("LH"),
		carriers.NewCarrier("LO"),
	}
	// expected result is sorted ascending
	expectedSegments := []segments.Segment{
		segments.NewSegment(0, 1, 1),
		segments.NewSegment(1, 2, 0),
	}

	// when
	b := segments.NewBuilder(airportList, carriers)
	b.Append("KKK", "ZZZ", "LH")
	b.Append("AAA", "KKK", "LO")
	actualSegments := b.Build()

	// then
	if len(expectedSegments) != len(actualSegments) {
		t.Fatalf("Num expected segments != num actual segments: %d : %d", len(expectedSegments), len(actualSegments))
	}

	for i, expected := range expectedSegments {
		actual := actualSegments[i]
		if actual != expected {
			t.Errorf("At index %d expected segment %v, got %v", i, expected, actual)
		}
	}
}
