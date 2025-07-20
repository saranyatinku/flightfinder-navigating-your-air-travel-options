package postprocessing

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

type CSVSegmentsToSegmentsFilter struct {
	airports airports.Airports
	carriers carriers.Carriers
}

func NewCSVSegmentsToSegmentsFilter(airports airports.Airports, carriers carriers.Carriers) *CSVSegmentsToSegmentsFilter {
	return &CSVSegmentsToSegmentsFilter{airports, carriers}
}

func (f *CSVSegmentsToSegmentsFilter) Filter(segs <-chan loading.CSVSegment) segments.Segments {
	sb := segments.NewBuilder(f.airports, f.carriers)

	for s := range segs {
		sb.Append(s.FromAirportCode, s.ToAirportCode, s.CarrierCode)
	}

	return sb.Build()
}
