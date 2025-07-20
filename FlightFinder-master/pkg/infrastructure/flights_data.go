package infrastructure

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
)

// FlightsData is just a flights data pack
type FlightsData struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Nations  nations.Nations
	Segments segments.Segments
}
