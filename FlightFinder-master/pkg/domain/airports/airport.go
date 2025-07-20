package airports

import "github.com/mateuszmidor/FlightFinder/pkg/domain/geo"

type ID int

const NullID = ID(-1)

type Airport struct {
	code      string // KRK
	name      string // Balice Krakow Airport
	nation    string // PL
	longitude geo.Longitude
	latitude  geo.Latitude
}

func NewAirport(code, name, nation string, longitude geo.Longitude, latitude geo.Latitude) Airport {
	return Airport{code, name, nation, longitude, latitude}
}

func NewAirportCodeOnly(code string) Airport {
	return Airport{code, "", "", 0, 0}
}

func (a *Airport) Code() string {
	return a.code
}

func (a *Airport) Name() string {
	return a.name
}

func (a *Airport) Nation() string {
	return a.nation
}
func (a *Airport) SetName(name string) {
	a.name = name
}

func (a *Airport) SetNation(nation string) {
	a.nation = nation
}

func (a *Airport) SetCoordinates(lon geo.Longitude, lat geo.Latitude) {
	a.longitude = lon
	a.latitude = lat
}

func (a *Airport) Longitude() geo.Longitude {
	return a.longitude
}

func (a *Airport) Latitude() geo.Latitude {
	return a.latitude
}
