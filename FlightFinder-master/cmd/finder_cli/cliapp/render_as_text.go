package cliapp

import (
	"fmt"
	"io"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

type PathRendererAsText struct {
	writer    io.Writer
	separator string
}

func NewPathRendererAsText(w io.Writer, pathSeparator string) *PathRendererAsText {
	return &PathRendererAsText{writer: w, separator: pathSeparator}
}

func (r *PathRendererAsText) Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData) {
	for nPath, path := range paths {
		renderPath(path, flightsData, r.writer)
		if nPath < len(paths)-1 {
			fmt.Fprint(r.writer, r.separator)
		}
	}
}

func renderPath(path pathfinding.Path, flightsData *infrastructure.FlightsData, writer io.Writer) {
	s0 := flightsData.Segments[path[0]]
	writer.Write([]byte(flightsData.Airports[s0.From()].Name()))
	for _, sID := range path {
		s := flightsData.Segments[sID]
		carrier := flightsData.Carriers[s.Carrier()].Code()
		toAirport := flightsData.Airports[s.To()].Name()
		fmt.Fprintf(writer, "-(%s)-%s", carrier, toAirport)
	}
}
