package apiserver

import (
	"encoding/json"
	"io"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

type PathRendererAsJSON struct {
	writer io.Writer
}

func NewPathRendererAsJSON(w io.Writer) *PathRendererAsJSON {
	return &PathRendererAsJSON{writer: w}
}

func (r *PathRendererAsJSON) Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData) {
	connections := make([]Connection, len((paths)))
	for nConnection, path := range paths {
		segment0 := flightsData.Segments[path[0]]
		airport0 := flightsData.Airports[segment0.From()]
		connections[nConnection].FromAirport = makeAirport(airport0, flightsData.Nations)
		connections[nConnection].Segments = makeSegments(path, flightsData)
	}
	json.NewEncoder(r.writer).Encode(connections)
}

func makeSegments(path pathfinding.Path, flightsData *infrastructure.FlightsData) []Segment {
	segments := make([]Segment, len(path))
	for nSegment, sID := range path {
		segment := flightsData.Segments[sID]
		segments[nSegment].Carrier = Carrier{Code: flightsData.Carriers[segment.Carrier()].Code()}
		segments[nSegment].ToAirport = makeAirport(flightsData.Airports[segment.To()], flightsData.Nations)
	}
	return segments
}

func makeAirport(a airports.Airport, nations nations.Nations) Airport {
	return Airport{
		Code:           a.Code(),
		Name:           a.Name(),
		Nation:         a.Nation(),
		NationFullName: nations[nations.GetByCode(a.Nation())].Name(),
		Lon:            float32(a.Longitude()),
		Lat:            float32(a.Latitude()),
	}
}
