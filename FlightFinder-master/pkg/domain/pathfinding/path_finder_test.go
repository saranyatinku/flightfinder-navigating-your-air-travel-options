package pathfinding

import (
	"strings"
	"testing"
)

// test helpful tools
const pathSeparator = ", "

func pathToString(path Path, connections exampleConnections) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	result := connections.nodes[connections.connections[path[0]].from].label + "-" +
		connections.connections[path[0]].label + "-" +
		connections.nodes[connections.connections[path[0]].to].label

	for i := 1; i < len(path); i++ {
		result += "-" + connections.connections[path[i]].label + "-" + connections.nodes[connections.connections[path[i]].to].label
	}

	return result
}

func pathsToString(paths []Path, connections exampleConnections) string {
	if len(paths) == 0 {
		return ""
	}

	var result string
	for _, p := range paths {
		result += pathToString(p, connections) + pathSeparator
	}

	return result
}

func cutOutSubstring(src, sub string) string {
	i := strings.Index(src, sub)
	if i == -1 {
		return src
	}

	return src[:i] + src[i+len(sub):]
}

func checkExpectedPaths(expected []string, actual []Path, connections exampleConnections, t *testing.T) {
	pathsString := pathsToString(actual, connections)
	for _, p := range expected {
		i := strings.Index(pathsString, p)
		if i == -1 {
			t.Errorf("Missing path: %s. Found: %s", p, pathsString)
		} else {
			pathsString = cutOutSubstring(pathsString, p+pathSeparator)
		}
	}
	if len(pathsString) > 0 {
		t.Errorf("Unexpected paths: %s", pathsString)
	}
}

/*
The general direction is: GDA -> KRK
Loops:
-RZE-WAW-KRK-RZE
-GDA-WAW-BYK-GDA
                          +-----+
                          | GDA |<------
                       /--+-----+\      \------------ (BY)
             (LH) /-----    /      \                  \-----
+-----+    /-----         / (LO)    \ (LO)              +-----+
| SZC | <--              v          \                   | BYK |
+-----+              +-----+         |                  +-----+
   \                 | BDG |         \                    -> ^
    \                +-----+          \        (LO)   ---/  /
     |                  /              \          ---/      |
     \                /-                \     ---/         /
      \              /                   v  -/             |
  (LH) \           /- (LO)            +-----+             /
        |         /                   | WAW |             |
        \        /                    +-----+<           / (BY)
         \     /-                        /    \          |
          v   <                         /      \ (LO)    /
        +-----+                        |        \-      |
        | WRO |                        /          \    /
        +-----+--\                    / (LO,BY)    \   |
                  ---\               /              \ /
                      ---\          v             +-----+
                    (LH)  -->+-----+       ------>| RZE |
                             | KRK |------/ (BY)  +-----+
                             +-----+
*/
type Map struct {
	connections exampleConnections
	// costs                                       Costs
	lub, gda, szc, bdg, byk, waw, wro, rze, krk NodeID
}

func makeFlightMap() Map {
	connections := exampleConnections{}
	// krk := AppendCity(&connections, City{"KRK", 0, 0})    // 0
	// rze := AppendCity(&connections, City{"RZE", 20, 5})   // 1
	// wro := AppendCity(&connections, City{"WRO", -20, 10}) // 2
	// waw := AppendCity(&connections, City{"WAW", 10, 20})  // 3
	// byk := AppendCity(&connections, City{"BYK", 30, 35})  // 4
	// bdg := AppendCity(&connections, City{"BDG", -10, 30}) // 5
	// szc := AppendCity(&connections, City{"SZC", -30, 30}) // 6
	// gda := AppendCity(&connections, City{"GDA", 0, 40})   // 7
	// lub := AppendCity(&connections, City{"LUB", 25, 10})  // 8, no connections

	krk := connections.addNode("KRK") // 0
	rze := connections.addNode("RZE") // 1
	wro := connections.addNode("WRO") // 2
	waw := connections.addNode("WAW") // 3
	byk := connections.addNode("BYK") // 4
	bdg := connections.addNode("BDG") // 5
	szc := connections.addNode("SZC") // 6
	gda := connections.addNode("GDA") // 7
	lub := connections.addNode("LUB") // 8
	// costs := Costs(nil)               // costs not used in multipath algo //NewCityTravelCosts(connections)

	connections.connect(gda, szc, "(LH)")
	connections.connect(szc, wro, "(LH)")
	connections.connect(gda, bdg, "(LO)")
	connections.connect(bdg, wro, "(LO)")
	connections.connect(wro, krk, "(LH)")
	connections.connect(gda, waw, "(LO)")
	connections.connect(waw, krk, "(LO)")
	connections.connect(waw, krk, "(BY)") // alt. BY connection
	connections.connect(krk, rze, "(BY)")
	connections.connect(rze, waw, "(LO)")
	connections.connect(waw, byk, "(LO)")
	connections.connect(byk, gda, "(BY)")
	connections.connect(rze, byk, "(BY)")
	connections.sort()

	return Map{
		connections: connections,
		// costs:       costs,
		lub: lub, gda: gda, szc: szc, bdg: bdg, byk: byk, waw: waw, wro: wro, rze: rze, krk: krk,
	}
}

func TestShouldFindSingleSegmentPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.byk, m.gda, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"BYK-(BY)-GDA",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindSingleAlternativeSegmentPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.wro, m.krk, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"WRO-(LH)-KRK",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindTwoSegmentsPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.byk, m.szc, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"BYK-(BY)-GDA-(LH)-SZC",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindTwoSegments2Paths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.gda, m.wro, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"GDA-(LO)-BDG-(LO)-WRO",
		"GDA-(LH)-SZC-(LH)-WRO",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindDirectAndIndirectPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.rze, m.waw, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"RZE-(LO)-WAW",
		"RZE-(BY)-BYK-(BY)-GDA-(LO)-WAW",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindIfSplittingPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.gda, m.krk, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"GDA-(LO)-WAW-(LO)-KRK",
		"GDA-(LO)-WAW-(BY)-KRK", // alt. BY connection
		"GDA-(LO)-BDG-(LO)-WRO-(LH)-KRK",
		"GDA-(LH)-SZC-(LH)-WRO-(LH)-KRK",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindIfMergingPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.byk, m.wro, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"BYK-(BY)-GDA-(LH)-SZC-(LH)-WRO",
		"BYK-(BY)-GDA-(LO)-BDG-(LO)-WRO",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldFindMultiLongPaths(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.gda, m.byk, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"GDA-(LO)-WAW-(LO)-BYK",
		"GDA-(LO)-WAW-(LO)-KRK-(BY)-RZE-(BY)-BYK",
		"GDA-(LO)-WAW-(BY)-KRK-(BY)-RZE-(BY)-BYK", // alt. BY connection
		"GDA-(LH)-SZC-(LH)-WRO-(LH)-KRK-(BY)-RZE-(BY)-BYK",
		"GDA-(LO)-BDG-(LO)-WRO-(LH)-KRK-(BY)-RZE-(BY)-BYK",
		"GDA-(LH)-SZC-(LH)-WRO-(LH)-KRK-(BY)-RZE-(LO)-WAW-(LO)-BYK",
		"GDA-(LO)-BDG-(LO)-WRO-(LH)-KRK-(BY)-RZE-(LO)-WAW-(LO)-BYK",
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func BenchmarkFindMultiLongPaths(b *testing.B) {
	m := makeFlightMap()
	for n := 0; n < b.N; n++ {
		_ = FindPaths(m.gda, m.byk, &m.connections, AlwaysContinueBuildingPaths)
	}
}

func TestShouldHandleCycle(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.waw, m.szc, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		"WAW-(LO)-BYK-(BY)-GDA-(LH)-SZC",
		"WAW-(LO)-KRK-(BY)-RZE-(BY)-BYK-(BY)-GDA-(LH)-SZC",
		"WAW-(BY)-KRK-(BY)-RZE-(BY)-BYK-(BY)-GDA-(LH)-SZC", // alt. BY connection
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}

func TestShouldHandleNoSuchPath(t *testing.T) {
	// given
	m := makeFlightMap()

	// when
	paths := FindPaths(m.waw, m.lub, &m.connections, AlwaysContinueBuildingPaths)

	// then
	expected := []string{
		// no such connection
	}
	checkExpectedPaths(expected, paths, m.connections, t)
}
