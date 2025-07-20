package application

import (
	"errors"
	"fmt"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

var NotFoundError = errors.New("not found error")

// AirportFinder allows for various searches in airport list
type AirportFinder struct {
	flightsData infrastructure.FlightsData
}

func NewAirportFinder(repo infrastructure.FlightsDataRepo) *AirportFinder {
	flightsData := repo.Load()
	return &AirportFinder{flightsData: flightsData}
}

func (a *AirportFinder) ByIATACode(code string) (airports.Airport, error) {
	id := a.flightsData.Airports.GetByCode(code)
	if id == airports.NullID {
		return airports.Airport{}, fmt.Errorf("airport IATA code %q not found: %w", code, NotFoundError)
	}
	return a.flightsData.Airports.Get(id), nil
}

func (a *AirportFinder) AllAirports() airports.Airports {
	return a.flightsData.Airports
}

func (a *AirportFinder) AirportsByCountry(twoLettersCode string) airports.Airports {
	result := make(airports.Airports, 0)
	for _, airport := range a.flightsData.Airports {
		if airport.Name() == twoLettersCode {
			result = append(result, airport)
		}
	}
	return result
}
