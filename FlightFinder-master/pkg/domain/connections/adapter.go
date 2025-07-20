package connections

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
)

// Adapter is pathfinding.Connections adapter for segments
type Adapter struct {
	segments segments.Segments // assumption: segments are sorted ascending by from, to
	finder   SegmentRangeFinder
}

// NewAdapter is constructor
func NewAdapter(segments segments.Segments) *Adapter {
	return &Adapter{segments, SegmentRangeFinder{}}
}

// GetDestinationNode implements pathfinding.Connections interface
func (a *Adapter) GetDestinationNode(connection pathfinding.ConnectionID) pathfinding.NodeID {
	return pathfinding.NodeID(a.segments[connection].To())
}

// GetOutgoingConnections implements pathfinding.Connections interface
func (a *Adapter) GetOutgoingConnections(node pathfinding.NodeID) (first, last pathfinding.ConnectionID) {
	f, l := a.finder.ByFromAirport(a.segments, airports.ID(node))
	return pathfinding.ConnectionID(f), pathfinding.ConnectionID(l)
}
