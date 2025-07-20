package loading

// CSVNation is nation in its denormalized form
type CSVNation struct {
	Code     string // eg. "PL"
	Iso      string // eg. "POL"
	Currency string // eg. "PLN"
	Name     string // eg. "POLAND"
}

// NewCSVNation is constructor
func NewCSVNation(code, iso, currency, name string) CSVNation {
	return CSVNation{
		Code:     code,
		Iso:      iso,
		Currency: currency,
		Name:     name,
	}
}
