package loading

import (
	"compress/gzip"
	"log"
	"os"
)

// StartLoadingSegmentsFromGzipCSV begins loading raw segments from gzipped csv into output channel
func StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV string, outSegments chan<- CSVSegment) {
	inputFile, err := os.Open(segmentsGzipCSV)
	if err != nil {
		log.Printf("Error opening %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		log.Printf("Error creating GZIP reader %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer gzipReader.Close()

	var loader SegmentLoader
	loader.StartLoading(gzipReader, outSegments)
}

// StartLoadingAirportsFromGzipCSV begins loading raw airports from gzipped csv into output channel
func StartLoadingAirportsFromGzipCSV(airportsGzipCSV string, outAirports chan<- CSVAirport) {
	inputFile, err := os.Open(airportsGzipCSV)
	if err != nil {
		log.Printf("Error opening %s: %v\n", airportsGzipCSV, err)
		close(outAirports)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		log.Printf("Error creating GZIP reader %s: %v\n", airportsGzipCSV, err)
		close(outAirports)
		return
	}
	defer gzipReader.Close()

	var loader AirportsLoader
	loader.StartLoading(gzipReader, outAirports)
}

// StartLoadingNationsFromGzipCSV begins loading raw nations from gzipped csv into output channel
func StartLoadingNationsFromGzipCSV(nationsGzipCSV string, outNations chan<- CSVNation) {
	inputFile, err := os.Open(nationsGzipCSV)
	if err != nil {
		log.Printf("Error opening %s: %v\n", nationsGzipCSV, err)
		close(outNations)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		log.Printf("Error creating GZIP reader %s: %v\n", nationsGzipCSV, err)
		close(outNations)
		return
	}
	defer gzipReader.Close()

	var loader NationsLoader
	loader.StartLoading(gzipReader, outNations)
}
