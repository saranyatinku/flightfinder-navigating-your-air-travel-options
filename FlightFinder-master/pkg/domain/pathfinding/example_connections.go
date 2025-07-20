package pathfinding

// This module is for testing purposes

import "sort"

type exampleNode struct {
	// longitude
	// latitude
	label string
}

type exampleConnection struct {
	from  NodeID
	to    NodeID
	label string
}

type exampleConnections struct {
	nodes       []exampleNode
	connections []exampleConnection
}

// Connections interface
func (e *exampleConnections) GetDestinationNode(connection ConnectionID) NodeID {
	return e.connections[connection].to
}

// Connections interface
func (e *exampleConnections) GetOutgoingConnections(node NodeID) (first ConnectionID, last ConnectionID) {
	return equalRange(e.connections, node)
}

//
// FROM NOW ON - HELPER FUNCTIONS
//

func (e *exampleConnections) addNode(label string) NodeID {
	id := NodeID(len(e.nodes))
	e.nodes = append(e.nodes, exampleNode{label})
	return id
}

func (e *exampleConnections) connect(from NodeID, to NodeID, label string) ConnectionID {
	conn := exampleConnection{from, to, label}
	id := ConnectionID(len(e.connections))
	e.connections = append(e.connections, conn)
	return id
}

func (e *exampleConnections) sort() {
	less := func(i, j int) bool {
		if e.connections[i].from != e.connections[j].from {
			return e.connections[i].from < e.connections[j].from
		}
		return e.connections[i].to < e.connections[j].to
	}
	sort.Slice(e.connections, less)
}

func equalRange(connections []exampleConnection, node NodeID) (first ConnectionID, last ConnectionID) {
	first = ConnectionID(len(connections))
	last = ConnectionID(len(connections))

	for i := 0; i < len(connections); i++ {
		if connections[i].from == node {
			first = ConnectionID(i)
			break
		}
	}

	for i := len(connections) - 1; i >= 0; i-- {
		if connections[i].from == node {
			last = ConnectionID(i + 1)
			break
		}
	}

	return
}
