package connections

import (
	"sort"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
)

type SegmentRangeFinder struct {
}

func (s *SegmentRangeFinder) ByFromAirport(segs segments.Segments, id airports.ID) (first, last segments.ID) {
	lo, hi := lowerBound(segs, id), upperBound(segs, id)
	if lo == hi {
		return segments.NullID, segments.NullID
	}
	return segments.ID(lo), segments.ID(hi)
}

// lowerBound is index of first matching element
func lowerBound(segments segments.Segments, id airports.ID) int {
	ge := func(i int) bool {
		return segments[i].From() >= id
	}

	return sort.Search(len(segments), ge)
}

// upperBound is index of last matching element +1
func upperBound(segments segments.Segments, id airports.ID) int {
	g := func(i int) bool {
		return segments[i].From() > id
	}

	return sort.Search(len(segments), g)
}
