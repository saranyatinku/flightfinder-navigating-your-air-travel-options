package postprocessing

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

func ExtractCarriers(segments <-chan loading.CSVSegment) carriers.Carriers {
	codes := make(map[string]bool)
	for s := range segments {
		codes[s.CarrierCode] = true
	}

	var list []carriers.Carrier
	for code := range codes {
		list = append(list, carriers.NewCarrier(code))
	}

	return carriers.NewCarriers(list...)
}
