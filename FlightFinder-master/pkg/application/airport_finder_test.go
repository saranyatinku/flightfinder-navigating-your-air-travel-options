package application_test

import (
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/application"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/geo"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
	"github.com/stretchr/testify/assert"
)

var KRK airports.Airport = newPolishAirport("KRK")
var GDN airports.Airport = newPolishAirport("GDN")
var LIM airports.Airport = newPeruAirport("LIM")
var CUZ airports.Airport = newPeruAirport("CUZ")

func newPolishAirport(code string) airports.Airport {
	return airports.NewAirport(code, "PL", "", geo.Longitude(0), geo.Latitude(0))
}

func newPeruAirport(code string) airports.Airport {
	return airports.NewAirport(code, "PE", "", geo.Longitude(0), geo.Latitude(0))
}

type stubRepo struct{}

func (stubRepo) Load() infrastructure.FlightsData {
	airports := airports.NewAirports(CUZ, LIM, GDN, KRK)
	return infrastructure.FlightsData{Airports: airports}
}

func TestAllAirportsReturnsAll(t *testing.T) {
	// given
	expected := airports.Airports{KRK, GDN, CUZ, LIM}
	svc := application.NewAirportFinder(stubRepo{})

	// when
	actual := svc.AllAirports()

	// then
	assert.ElementsMatch(t, expected, actual)
}

func TestAirportsByCountryReturnsAllMatchingAirports(t *testing.T) {
	// given
	expectedPL := airports.Airports{KRK, GDN}
	expectedPE := airports.Airports{LIM, CUZ}
	svc := application.NewAirportFinder(stubRepo{})

	// when
	actualPL := svc.AirportsByCountry("PL")
	actualPE := svc.AirportsByCountry("PE")

	// then
	assert.ElementsMatch(t, expectedPL, actualPL)
	assert.ElementsMatch(t, expectedPE, actualPE)
}
