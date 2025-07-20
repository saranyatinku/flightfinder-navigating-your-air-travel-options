// +build ignore

package pathfinding

type CostType = float64
type Costs interface {
	// actual, known cost
	D(from NodeID, to NodeID) CostType

	// heuristic, not yet known cost
	H(current NodeID, goal NodeID) CostType
}
