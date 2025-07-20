package carriers

import "sort"

// Carriers holds mapping for id-carrier
type Carriers []Carrier

// NewCarriers creates Carriers list sorted ascending by code. This is used for binary search
func NewCarriers(carriers ...Carrier) Carriers {
	less := func(i, j int) bool {
		return carriers[i].code < carriers[j].code
	}

	sort.Slice(carriers, less)
	return carriers
}

// Get returns Carrier by CarrierID
func (c Carriers) Get(id ID) Carrier {
	return c[id]
}

// GetByCode returns CarrierID of given carrier
// Precondition: carriers are sorted ascending
func (c Carriers) GetByCode(code string) ID {
	ge := func(i int) bool {
		return c[i].code >= code
	}

	foundIndex := sort.Search(len(c), ge)

	if c[foundIndex].code != code {
		return NullID
	}

	return ID(foundIndex)
}
