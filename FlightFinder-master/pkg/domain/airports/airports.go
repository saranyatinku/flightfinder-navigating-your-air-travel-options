package airports

import "sort"

// Airports holds mapping for id-airport
type Airports []Airport

// NewAirports creates Airports list sorted ascending by code. This is used for binary search
func NewAirports(airports ...Airport) Airports {
	less := func(i, j int) bool {
		return airports[i].code < airports[j].code
	}

	sort.Slice(airports, less)
	return airports
}

// Get returns Airport by ID
func (a Airports) Get(id ID) Airport {
	return a[id]
}

// GetByCode returns ID of given airport
// Precondition: airports are sorted ascending by code
func (a Airports) GetByCode(code string) ID {
	ge := func(i int) bool {
		return a[i].code >= code
	}

	foundIndex := sort.Search(len(a), ge)

	if foundIndex < 0 || foundIndex >= len(a) || a[foundIndex].code != code {
		return NullID
	}

	return ID(foundIndex)
}
