package segments

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
)

// ID is segment uid
type ID int

// NullID means no such segment
const NullID = ID(-1)

// Segment is smallest part of the journey
type Segment struct {
	from    airports.ID
	to      airports.ID
	carrier carriers.ID
}

// NewSegment is constructor
func NewSegment(from, to airports.ID, carrier carriers.ID) Segment {
	return Segment{from, to, carrier}
}

// From tells where the journey begins
func (s *Segment) From() airports.ID {
	return s.from
}

// To tells where the journey ends
func (s *Segment) To() airports.ID {
	return s.to
}

// Carrier tells the carrier
func (s *Segment) Carrier() carriers.ID {
	return s.carrier
}
