package pathfinding

// CheckContinueBuildingPaths allows stop processing current path on demand
type CheckContinueBuildingPaths func(currentPathLen, totalPathsFound int) bool

// AlwaysContinueBuildingPaths never stops processing current path
func AlwaysContinueBuildingPaths(currentPathLen, totalPathsFound int) bool {
	return true
}

// FindPaths find all possible paths, not only the least costly
// calling continueBuildingPaths costs +40% finding time: 5->7sec. But enables extensibility :)
func FindPaths(start NodeID, goal NodeID, connections Connections, continueBuildingPaths CheckContinueBuildingPaths) (result []Path) {
	const maxNodeID = 1000000 // assumption: max nodeID in graph < 1'000'000; for airports should be enough
	var alreadyVisited [maxNodeID]bool
	var recursive func(current NodeID, pathLength int)
	currPath := make(Path, 10)

	recursive = func(current NodeID, pathLength int) {
		if pathLength >= len(currPath) || !continueBuildingPaths(pathLength, len(result)) {
			return
		}

		// check if path found
		if current == goal {
			newPath := make(Path, pathLength)
			copy(newPath, currPath)
			result = append(result, newPath)
			return
		}

		// mark current as visited. For cycle detection
		alreadyVisited[current] = true

		// check all outgoing connections of current
		first, last := connections.GetOutgoingConnections(current)
		for connID := first; connID < last; connID++ {
			dest := connections.GetDestinationNode(connID)

			// avoid cycles
			if alreadyVisited[dest] {
				continue
			}

			// process outgoing connection
			currPath[pathLength] = connID

			recursive(dest, pathLength+1)

			alreadyVisited[dest] = false // dest added by "recursive"
		}

	}

	recursive(start, 0)
	return result
}
