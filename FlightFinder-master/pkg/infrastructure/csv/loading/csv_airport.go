package loading

import "github.com/mateuszmidor/FlightFinder/pkg/domain/geo"

// CSVAirport is airport in its denormalized form
type CSVAirport struct {
	AirportCode string        // eg. "GDN"
	FullName    string        // eg. "Gdynia Airport"
	Nation      string        // eg. "PL"
	Longitude   geo.Longitude // eg. 20.1
	Latitude    geo.Latitude  // eg. -50.4
}

// NewCSVAirport is constructor
func NewCSVAirport(code, name, nation string, lng geo.Longitude, lat geo.Latitude) CSVAirport {
	return CSVAirport{
		AirportCode: code,
		FullName:    name,
		Nation:      nation,
		Longitude:   lng,
		Latitude:    lat,
	}
}
