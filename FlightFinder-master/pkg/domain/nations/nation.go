package nations

// ID uniquely identifies nation on the list
type ID int

// NullID is used to indicate that no such nation exists
const NullID = ID(-1)

// Nation describes country where an airport is located
type Nation struct {
	code     string // PL
	iso      string // POL
	currency string // PLN
	name     string // POLAND
}

// NewNation is constructor
func NewNation(code, iso, currency, name string) Nation {
	return Nation{
		code:     code,
		iso:      iso,
		currency: currency,
		name:     name,
	}
}

// Code returns 2 letter code eg PL
func (n *Nation) Code() string {
	return n.code
}

// Name returns full name eg POLAND
func (n *Nation) Name() string {
	return n.name
}

// Iso returns 3 letter country code eg POL
func (n *Nation) Iso() string {
	return n.iso
}

// Currency returns 3 letter curency eg PLN
func (n *Nation) Currency() string {
	return n.currency
}
