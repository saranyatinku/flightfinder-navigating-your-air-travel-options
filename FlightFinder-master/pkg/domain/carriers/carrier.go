package carriers

// ID is program internal carrier id
type ID int

const NullID = ID(-1)

// Carrier is company who provides transportation between 2 airports
type Carrier struct {
	code string // eg. "LO"
}

// NewCarrier is constructor
func NewCarrier(code string) Carrier {
	return Carrier{code}
}

// Code is two characters carrier code eg "LO"
func (c *Carrier) Code() string {
	return c.code
}
