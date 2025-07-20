package application

import (
	"fmt"
	"sort"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/connections"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/pathfinding"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
)

type PathRenderer interface {
	Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData)
}

type ConnectionFinder struct {
	flightsData infrastructure.FlightsData
	connections pathfinding.Connections
}

func NewConnectionFinder(repo infrastructure.FlightsDataRepo) *ConnectionFinder {

	flightsData := repo.Load()
	connections := connections.NewAdapter(flightsData.Segments)
	return &ConnectionFinder{flightsData: flightsData, connections: connections}
}

func (f *ConnectionFinder) Find(fromAirportCode, toAirportCode string, maxSegmentCount int, pathRenderer PathRenderer) error {
	fromID, toID, err := getFromToAirportID(fromAirportCode, toAirportCode, f.flightsData.Airports)
	if err != nil {
		return err
	}

	limiter := makeLimiter(maxSegmentCount)
	paths := pathfinding.FindPaths(pathfinding.NodeID(fromID), pathfinding.NodeID(toID), f.connections, limiter)
	sortPathsByNumSegmentsAscending(paths)

	pathRenderer.Render(paths, &f.flightsData)
	return nil
}

func getFromToAirportID(fromAirportCode, toAirportCode string, airprts airports.Airports) (airports.ID, airports.ID, error) {
	from := airprts.GetByCode(fromAirportCode)
	if from == airports.NullID {
		return airports.NullID, airports.NullID, fmt.Errorf(`Invalid "from" airport: %s`, fromAirportCode)
	}

	to := airprts.GetByCode(toAirportCode)
	if to == airports.NullID {
		return airports.NullID, airports.NullID, fmt.Errorf(`Invalid "to" airport: %s`, toAirportCode)
	}

	return from, to, nil
}

func sortPathsByNumSegmentsAscending(paths []pathfinding.Path) {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

func makeLimiter(maxSegmentCount int) pathfinding.CheckContinueBuildingPaths {
	return func(currentPathLen, totalPathsFound int) bool {
		maxPathLen := maxSegmentCount + 1 // KRK-WAW-GDN is 2 segments made of 3 airports
		return currentPathLen < maxPathLen && totalPathsFound < 1000
	}
}
