package cliapp_test

import (
	"bytes"
	"testing"

	"github.com/mateuszmidor/FlightFinder/cmd/finder_cli/cliapp"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/geo"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
	"github.com/stretchr/testify/assert"
)

const CarrierBlack = 000
const CarrierRed = 001
const CarrierGreen = 002
const CarrierBlue = 003
const CarrierFuchsia = 004

var flightData = infrastructure.FlightsData{
	Airports: airports.Airports{
		airports.NewAirport("0", "airport0", "nation0", geo.Longitude(0), geo.Latitude(0)),
		airports.NewAirport("1", "airport1", "nation1", geo.Longitude(0), geo.Latitude(0)),
		airports.NewAirport("2", "airport2", "nation2", geo.Longitude(0), geo.Latitude(0)),
		airports.NewAirport("3", "airport3", "nation3", geo.Longitude(0), geo.Latitude(0)),
		airports.NewAirport("4", "airport4", "nation3", geo.Longitude(0), geo.Latitude(0)),
	},
	Carriers: carriers.Carriers{
		carriers.NewCarrier("black"),
		carriers.NewCarrier("red"),
		carriers.NewCarrier("green"),
		carriers.NewCarrier("blue"),
		carriers.NewCarrier("fuchsia"),
	},
	Segments: segments.Segments{
		segments.NewSegment(0, 1, CarrierBlack),   // connectionID=0
		segments.NewSegment(1, 2, CarrierRed),     // connectionID=1
		segments.NewSegment(2, 3, CarrierGreen),   // connectionID=2
		segments.NewSegment(3, 4, CarrierBlue),    // connectionID=3
		segments.NewSegment(4, 1, CarrierFuchsia), // connectionID=4
	},
}

func TestPathRendererTurnsEmptyPathsIntoEmptyPathString(t *testing.T) {
	// given
	flightData := infrastructure.FlightsData{}
	paths := []pathfinding.Path{}
	expected := ""
	buf := bytes.NewBuffer([]byte{})
	pathRenderer := cliapp.NewPathRendererAsText(buf, "")

	// when
	pathRenderer.Render(paths, &flightData)
	actual := buf.String()

	// then
	assert.Equal(t, expected, actual, "Paths: %v", paths)
}

func TestPathRendererTurnsValidSinglePathIntoValidPathString(t *testing.T) {
	// given
	path := pathfinding.Path{
		pathfinding.ConnectionID(1), // connectionID=1
		pathfinding.ConnectionID(2), // connectionID=2
		pathfinding.ConnectionID(3), // connectionID=3
	}
	paths := []pathfinding.Path{path}
	expected := "airport1-(red)-airport2-(green)-airport3-(blue)-airport4"
	buf := bytes.NewBuffer([]byte{})
	pathRenderer := cliapp.NewPathRendererAsText(buf, "")

	// when
	pathRenderer.Render(paths, &flightData)
	actual := buf.String()

	// then
	assert.Equal(t, expected, actual, "Paths: %v", paths)
}

func TestPathRendererTurnsValidMultiplePathsIntoValidPathString(t *testing.T) {
	// given
	path1 := pathfinding.Path{
		pathfinding.ConnectionID(1), // connectionID=1
		pathfinding.ConnectionID(2), // connectionID=2
		pathfinding.ConnectionID(3), // connectionID=3
	}
	path2 := pathfinding.Path{
		pathfinding.ConnectionID(4), // connectionID=4
		pathfinding.ConnectionID(1), // connectionID=1
		pathfinding.ConnectionID(2), // connectionID=2
	}
	paths := []pathfinding.Path{path1, path2}
	expected := "airport1-(red)-airport2-(green)-airport3-(blue)-airport4,airport4-(fuchsia)-airport1-(red)-airport2-(green)-airport3"
	buf := bytes.NewBuffer([]byte{})
	pathRenderer := cliapp.NewPathRendererAsText(buf, ",")

	// when
	pathRenderer.Render([]pathfinding.Path{path1, path2}, &flightData)
	actual := buf.String()

	// then
	assert.Equal(t, expected, actual, "Paths: %v", paths)
}
