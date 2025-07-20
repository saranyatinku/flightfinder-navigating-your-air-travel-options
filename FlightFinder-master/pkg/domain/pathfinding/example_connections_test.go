package pathfinding

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	// given
	connections := &exampleConnections{}

	// when
	n0 := connections.addNode("n0")
	n1 := connections.addNode("n1")
	n2 := connections.addNode("n2")

	// then
	if n0 != 0 {
		t.Errorf("Expected NodeID = 0, got: %d", n0)
	}

	if n1 != 1 {
		t.Errorf("Expected NodeID = 1, got: %d", n1)
	}

	if n2 != 2 {
		t.Errorf("Expected NodeID = 2, got: %d", n2)
	}

	if label := connections.nodes[n0].label; label != "n0" {
		t.Errorf("Expected node label = n0, got: %s", label)
	}

	if label := connections.nodes[n1].label; label != "n1" {
		t.Errorf("Expected node label = n1, got: %s", label)
	}

	if label := connections.nodes[n2].label; label != "n2" {
		t.Errorf("Expected node label = n2, got: %s", label)
	}
}

func TestConnect(t *testing.T) {
	// given
	connections := &exampleConnections{}
	n0 := connections.addNode("n0")
	n1 := connections.addNode("n1")
	n2 := connections.addNode("n2")

	// when
	ca := connections.connect(n0, n1, "ca")
	cb := connections.connect(n1, n2, "cb")

	// then
	if ca != 0 {
		t.Errorf("Expected ConnectionID = 0, got: %d", ca)
	}

	if cb != 1 {
		t.Errorf("Expected ConnectionID = 1, got: %d", cb)
	}

	if label := connections.connections[ca].label; label != "ca" {
		t.Errorf("Expected connection label = ca, got: %s", label)
	}

	if label := connections.connections[cb].label; label != "cb" {
		t.Errorf("Expected connection label = cb, got: %s", label)
	}
}

func TestSort(t *testing.T) {
	// given
	connections := &exampleConnections{}
	connections.connect(2, 10, "2-10")
	connections.connect(2, 20, "2-20")
	connections.connect(3, 20, "3-20")
	connections.connect(3, 10, "3-10")
	connections.connect(1, 10, "1-10")
	connections.connect(1, 20, "1-20")

	// when
	connections.sort()

	// then
	if conn := connections.connections[0]; conn.label != "1-10" {
		t.Errorf("Expected connections[0] = 1-10, got %s", conn.label)
	}

	if conn := connections.connections[1]; conn.label != "1-20" {
		t.Errorf("Expected connections[1] = 1-20, got %s", conn.label)
	}

	if conn := connections.connections[2]; conn.label != "2-10" {
		t.Errorf("Expected connections[2] = 2-10, got %s", conn.label)
	}

	if conn := connections.connections[3]; conn.label != "2-20" {
		t.Errorf("Expected connections[3] = 2-20, got %s", conn.label)
	}

	if conn := connections.connections[4]; conn.label != "3-10" {
		t.Errorf("Expected connections[4] = 3-10, got %s", conn.label)
	}

	if conn := connections.connections[5]; conn.label != "3-20" {
		t.Errorf("Expected connections[5] = 3-20, got %s", conn.label)
	}
}

func TestEqualRange(t *testing.T) {
	// given
	connections := &exampleConnections{}
	connections.connect(1, 10, "1-10")
	connections.connect(1, 20, "1-20")
	connections.connect(2, 10, "2-10")
	connections.connect(2, 20, "2-20")
	connections.connect(2, 30, "2-30")
	connections.connect(3, 10, "3-10")
	connections.connect(3, 20, "3-20")
	connections.sort()

	cases := []struct {
		node        NodeID
		first, last ConnectionID
	}{
		{1, 0, 2},
		{2, 2, 5},
		{3, 5, 7},
		{4, 7, 7}, // no such nodeID; 7 is past the total connnection count
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Searching nodeid = %d", c.node), func(t *testing.T) {
			// when
			first, last := equalRange(connections.connections, c.node)

			// then
			if first != c.first {
				t.Errorf("Expected first = %d, got %d", c.first, first)
			}
			if last != c.last {
				t.Errorf("Expected last = %d, got %d", c.last, last)
			}
		})
	}

}

func TestGetDestinationNode(t *testing.T) {
	// given
	connections := &exampleConnections{}
	c1 := connections.connect(1, 10, "1-10")
	c2 := connections.connect(2, 20, "2-20")
	c3 := connections.connect(3, 30, "3-30")

	cases := []struct {
		c ConnectionID
		n NodeID
	}{
		{c1, 10},
		{c2, 20},
		{c3, 30},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Getting destination node for connectionID = %d", c.n), func(t *testing.T) {
			// when
			n := connections.GetDestinationNode(c.c)

			// then
			if n != c.n {
				t.Errorf("Expected nodeID = %d, got %d", c.n, n)
			}
		})
	}
}

func TestGetOutgoingConnections(t *testing.T) {
	// given
	connections := &exampleConnections{}
	connections.connect(1, 2, "1-2") // 0
	connections.connect(2, 3, "2-3") // 1
	connections.connect(3, 5, "3-5") // 2
	connections.connect(3, 4, "3-4") // 3
	connections.sort()

	cases := []struct {
		node        NodeID
		first, last ConnectionID
	}{
		{1, 0, 1},
		{2, 1, 2},
		{3, 2, 4},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Getting outgoing connections for nodeID = %d", c.node), func(t *testing.T) {
			// when
			first, last := connections.GetOutgoingConnections(c.node)

			// then
			if first != c.first {
				t.Errorf("Expected first = %d, got %d", c.first, first)
			}
			if last != c.last {
				t.Errorf("Expected last = %d, got %d", c.last, last)
			}
		})
	}
}
