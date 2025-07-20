package connections_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/connections"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
)

func TestFindByOriginReturnsProperRange(t *testing.T) {
	// given
	segmentList := segments.Segments{
		segments.NewSegment(1, 10, -1),
		segments.NewSegment(1, 20, -1),
		segments.NewSegment(2, 10, -1),
		segments.NewSegment(2, 20, -1),
		segments.NewSegment(2, 30, -1),
		segments.NewSegment(4, 10, -1),
		segments.NewSegment(4, 20, -1),
	}

	cases := []struct {
		id          airports.ID
		first, last segments.ID
	}{
		{1, 0, 2},
		{2, 2, 5},
		{3, segments.NullID, segments.NullID}, // no such from airport
		{4, 5, 7},
		{5, segments.NullID, segments.NullID}, // no such from airport
	}

	var finder connections.SegmentRangeFinder
	for _, c := range cases {
		t.Run(fmt.Sprintf("Searching ID = %d", c.id), func(t *testing.T) {
			// when
			first, last := finder.ByFromAirport(segmentList, c.id)

			// then
			if first != c.first {
				t.Errorf("Expected first = %d, got %d", c.first, first)
			}
			if last != c.last {
				t.Errorf("Expected last = %d, got %d", c.last, last)
			}
		})
	}
}
