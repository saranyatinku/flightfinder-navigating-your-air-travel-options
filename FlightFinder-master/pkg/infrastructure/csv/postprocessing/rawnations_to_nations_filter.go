package postprocessing

import (
	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

// FilterNations turns stream of CSVNation into Nations list
func FilterNations(rawnations <-chan loading.CSVNation) nations.Nations {
	var list []nations.Nation
	for n := range rawnations {
		list = append(list, nations.NewNation(n.Code, n.Iso, n.Currency, n.Name))
	}

	return nations.NewNations(list...)
}
