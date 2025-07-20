package pathfinding

// Connections represents a network of interconnected nodes for the sake of path finding
// first is first outgoing connectionID
// last is last outgoing connectionID +1
// , so to iterate over connections, do: for i := first; i < last; i++
type Connections interface {
	GetOutgoingConnections(node NodeID) (first ConnectionID, last ConnectionID)
	GetDestinationNode(connection ConnectionID) NodeID
}
