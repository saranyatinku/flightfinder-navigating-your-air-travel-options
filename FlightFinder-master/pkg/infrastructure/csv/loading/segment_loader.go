package loading

import (
	"encoding/csv"
	"io"
)

const numSegmentCSVColumns = 3

// SegmentLoader loads segments from given source
type SegmentLoader struct {
}

// StartLoading starts loading raw segments into output channel
// Pipeline instead batch load approach to accomodate segment database that would exceed machine ram limitations
// Usage: go source.StartLoading(...)
func (r *SegmentLoader) StartLoading(reader io.Reader, outputSegments chan<- CSVSegment) {
	csv := csv.NewReader(reader)
	csv.ReuseRecord = true
	csv.FieldsPerRecord = numSegmentCSVColumns

	for {
		rec, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err == nil && rec != nil {
			outputSegments <- parseCSVSegment(rec)
		}

	}

	close(outputSegments)
}

func parseCSVSegment(data []string) CSVSegment {
	// CSV structure:
	// "KRK","KTW","LO"

	return CSVSegment{
		FromAirportCode: data[0],
		ToAirportCode:   data[1],
		CarrierCode:     data[2],
	}
}
