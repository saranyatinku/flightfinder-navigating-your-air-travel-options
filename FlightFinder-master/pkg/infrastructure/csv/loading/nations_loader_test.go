package loading_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
)

func TestLoadValidCSVShouldReturnAllNations(t *testing.T) {
	// given
	var loader loading.NationsLoader
	actualNations := make(chan loading.CSVNation, 1)
	expectedNations := []loading.CSVNation{
		loading.NewCSVNation("CN", "CHN", "CNY", "CHINA"),
		loading.NewCSVNation("ES", "ESP", "EUR", "SPAIN"),
		loading.NewCSVNation("PL", "POL", "PLN", "POLAND"),
	}
	// NATION,ISO,CURRENCY,DESCRIPTION
	csv := `
"CN","CHN","CNY","CHINA"
"ES","ESP","EUR","SPAIN"
"PL","POL","PLN","POLAND"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualNations)

	// then
	errorDetails := checkExpectedNations(expectedNations, actualNations)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedNations(expected []loading.CSVNation, actual chan loading.CSVNation) string {
	var result string
	for seg := range actual {
		if index := findNation(seg, expected); index != -1 {
			removeNation(index, &expected)
		} else {
			result += fmt.Sprintf("Unexpected Nation loaded: %+v\n", seg)
		}
	}

	if len(expected) != 0 {
		result += fmt.Sprintf("Expected Nations not loaded: %+v", expected)
	}

	return result
}

func findNation(subject loading.CSVNation, list []loading.CSVNation) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeNation(index int, list *[]loading.CSVNation) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
