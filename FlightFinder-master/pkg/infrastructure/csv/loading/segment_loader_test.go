package loading_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

func TestLoadValidCSVShouldReturnAllSegments(t *testing.T) {
	// given
	var loader loading.SegmentLoader
	actualSegments := make(chan loading.CSVSegment, 1)
	expectedSegments := []loading.CSVSegment{
		{FromAirportCode: "GDY", ToAirportCode: "WAW", CarrierCode: "BY"},
		{FromAirportCode: "WAW", ToAirportCode: "KRK", CarrierCode: "LH"},
		{FromAirportCode: "KRK", ToAirportCode: "KTW", CarrierCode: "LO"},
	}
	csv := `
"GDY","WAW","BY"
"WAW","KRK","LH"
"KRK","KTW","LO"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func TestLoadBrokenCSVShouldReturnOnlyValidSegments(t *testing.T) {
	// given
	var loader loading.SegmentLoader
	actualSegments := make(chan loading.CSVSegment, 1)
	expectedSegments := []loading.CSVSegment{
		// {"GDY", "WAW", "BY"},
		// {"WAW", "KRK", "LH"},
		{"KRK", "KTW", "LO"},
	}
	csv := `
"GDY","WAW", "BY"
"WAW","KRK"
"KRK","KTW","LO"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedSegments(expected []loading.CSVSegment, actual chan loading.CSVSegment) string {
	var result string
	for seg := range actual {
		if index := findSegment(seg, expected); index != -1 {
			removeSegment(index, &expected)
		} else {
			result += fmt.Sprintf("Unexpected segment loaded: %v\n", seg)
		}
	}

	if len(expected) != 0 {
		result += fmt.Sprintf("Expected segments not loaded: %v", expected)
	}

	return result
}

func findSegment(subject loading.CSVSegment, list []loading.CSVSegment) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeSegment(index int, list *[]loading.CSVSegment) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
