package apiserver_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	apiserver "github.com/mateuszmidor/FlightFinder/cmd/finder_web/webapp/apiserver/go"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/geo"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathRendererTurnsEmptyPathsIntoEmptyJSON(t *testing.T) {
	// given
	flightData := infrastructure.FlightsData{}
	paths := []pathfinding.Path{}
	expected := "[]"
	buf := bytes.NewBuffer([]byte{})
	renderer := apiserver.NewPathRendererAsJSON(buf)

	// when
	renderer.Render(paths, &flightData)

	// then
	assert.Equal(t, expected, strings.TrimSpace(buf.String()))
}

func TestPathRendererTurnsValidPathIntoValidPathJson(t *testing.T) {
	// given
	flightsData := infrastructure.FlightsData{
		Airports: airports.NewAirports(
			airports.NewAirport("GDN", "GDANSK", "PL", geo.Longitude(51), geo.Latitude(21)),
			airports.NewAirport("BCN", "BARCELONA", "ES", geo.Longitude(60), geo.Latitude(10)),
			airports.NewAirport("KRK", "KRAKOW", "PL", geo.Longitude(49), geo.Latitude(19)),
		),
		Carriers: carriers.Carriers{
			carriers.NewCarrier("LO"),
			carriers.NewCarrier("FR"),
		},
		Nations: nations.Nations{
			nations.NewNation("ES", "-", "-", "SPAIN"),
			nations.NewNation("PL", "-", "-", "POLAND"),
		},
		// notice: segments must be sorted by from airport
		Segments: segments.Segments{
			segments.NewSegment(2, 0, 0), // connectionID=0 : KRK-BCN
			segments.NewSegment(0, 1, 1), // connectionID=1 : BCN-GDN
		},
	}
	path := pathfinding.Path{
		pathfinding.ConnectionID(0),
		pathfinding.ConnectionID(1),
	}
	paths := []pathfinding.Path{path}

	// KRK-(LO)-WAW-(FR)-GDN
	expected := apiserver.Connection{
		FromAirport: apiserver.Airport{
			Code:           "KRK",
			Name:           "KRAKOW",
			Nation:         "PL",
			NationFullName: "POLAND",
			Lon:            49.0,
			Lat:            19.0,
		},
		Segments: []apiserver.Segment{
			{
				Carrier: apiserver.Carrier{
					Code: "LO",
				},
				ToAirport: apiserver.Airport{
					Code:           "BCN",
					Name:           "BARCELONA",
					Nation:         "ES",
					NationFullName: "SPAIN",
					Lon:            60.0,
					Lat:            10.0,
				},
			},
			{
				Carrier: apiserver.Carrier{
					Code: "FR",
				},
				ToAirport: apiserver.Airport{
					Code:           "GDN",
					Name:           "GDANSK",
					Nation:         "PL",
					NationFullName: "POLAND",
					Lon:            51.0,
					Lat:            21.0,
				},
			},
		},
	}
	buf := bytes.NewBuffer([]byte{})
	renderer := apiserver.NewPathRendererAsJSON(buf)

	// when
	renderer.Render(paths, &flightsData)

	// then
	var actualPaths []apiserver.Connection
	json.NewDecoder(buf).Decode(&actualPaths)
	require.Len(t, actualPaths, 1, "For single input path there should be single path outputted in json, got %d", len(actualPaths))

	actual := actualPaths[0]
	assert.Equal(t, expected.FromAirport, actual.FromAirport)

	require.Equal(t, len(actual.Segments), len(expected.Segments))
	assert.Equal(t, actual.Segments[0], expected.Segments[0])
	assert.Equal(t, actual.Segments[1], expected.Segments[1])
}
