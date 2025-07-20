package csv

import (
	"path"

	"github.com/mateuszmidor/FlightFinder/pkg/domain/airports"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/carriers"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/nations"
	"github.com/mateuszmidor/FlightFinder/pkg/domain/segments"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/loading"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv/postprocessing"
)

// FlightsDataRepoCSV implements FlightsDataRepo
type FlightsDataRepoCSV struct {
	csvFilesDirectory string
}

// NewFlightsDataRepoCSV is constructor
func NewFlightsDataRepoCSV(csvFilesDirectory string) *FlightsDataRepoCSV {
	return &FlightsDataRepoCSV{csvFilesDirectory: csvFilesDirectory}
}

// Load reads CSV files and returns flights data
func (r *FlightsDataRepoCSV) Load() infrastructure.FlightsData {
	nations := loadNationsFromCSV(r.csvFilesDirectory)
	airports := loadAirportsFromCSV(r.csvFilesDirectory)
	carriers := loadCarriersFromCSV(r.csvFilesDirectory)
	segments := loadSegmentsFromCSV(airports, carriers, r.csvFilesDirectory)

	return infrastructure.FlightsData{
		Nations:  nations,
		Airports: airports,
		Carriers: carriers,
		Segments: segments,
	}
}

func loadNationsFromCSV(csvDataDirectory string) nations.Nations {
	nationsFilePath := getNationsCSVFilePath(csvDataDirectory)

	// load nations from nations file
	csvNations := make(chan loading.CSVNation, 100)
	go loading.StartLoadingNationsFromGzipCSV(nationsFilePath, csvNations)
	return postprocessing.FilterNations(csvNations)
}

func loadCarriersFromCSV(csvDataDirectory string) carriers.Carriers {
	segmentsFilePath := getSegmentsCSVFilePath(csvDataDirectory)

	// extract carriers from segments
	csvSegments := make(chan loading.CSVSegment, 100)
	go loading.StartLoadingSegmentsFromGzipCSV(segmentsFilePath, csvSegments)
	carriersUsedBySegments := postprocessing.ExtractCarriers(csvSegments)
	return carriersUsedBySegments
}

func loadAirportsFromCSV(csvDataDirectory string) airports.Airports {
	segmentsFilePath := getSegmentsCSVFilePath(csvDataDirectory)
	airportsFilePath := getAirportsCSVFilePath(csvDataDirectory)

	// extract airports from segments
	csvSegments := make(chan loading.CSVSegment, 100)
	go loading.StartLoadingSegmentsFromGzipCSV(segmentsFilePath, csvSegments)
	airportsUsedBySegments := postprocessing.ExtractAirports(csvSegments)

	// enrich extracted airports with additional data from airports file
	csvAirports := make(chan loading.CSVAirport, 100)
	go loading.StartLoadingAirportsFromGzipCSV(airportsFilePath, csvAirports)
	postprocessing.EnrichAirports(airportsUsedBySegments, csvAirports)
	return airportsUsedBySegments
}

func loadSegmentsFromCSV(a airports.Airports, c carriers.Carriers, csvDataDirectory string) segments.Segments {
	segmentsFilePath := getSegmentsCSVFilePath(csvDataDirectory)

	csvSegments := make(chan loading.CSVSegment, 100)
	go loading.StartLoadingSegmentsFromGzipCSV(segmentsFilePath, csvSegments)
	return postprocessing.NewCSVSegmentsToSegmentsFilter(a, c).Filter(csvSegments)
}

func getNationsCSVFilePath(csvDataDirectory string) string {
	const nationsFileName = "nations.csv.gz"
	return path.Join(csvDataDirectory, nationsFileName)
}

func getAirportsCSVFilePath(csvDataDirectory string) string {
	const airportsFileName = "airports.csv.gz"
	return path.Join(csvDataDirectory, airportsFileName)
}

func getSegmentsCSVFilePath(csvDataDirectory string) string {
	const segmentsFileName = "segments.csv.gz"
	return path.Join(csvDataDirectory, segmentsFileName)
}
