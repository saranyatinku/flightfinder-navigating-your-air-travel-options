package infrastructure

// FlightsDataRepo is interface for loading flights data
type FlightsDataRepo interface {
	Load() FlightsData
}
