package loading

// CSVSegment is segment in its denormalized form
type CSVSegment struct {
	FromAirportCode string // eg. "GDN"
	ToAirportCode   string // eg. "KRK"
	CarrierCode     string // eg. "LO"
}

// NewCSVSegment is constructor
func NewCSVSegment(from, to, carrier string) CSVSegment {
	return CSVSegment{
		FromAirportCode: from,
		ToAirportCode:   to,
		CarrierCode:     carrier,
	}
}
