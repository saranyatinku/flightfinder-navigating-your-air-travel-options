package loading

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

const numNationCSVColumns = 4

// NationsLoader loads nations from given source
type NationsLoader struct {
}

// StartLoading starts loading raw nations into output channel
// Pipeline instead batch load approach to save memory
// Usage: go source.StartLoading(...)
func (r *NationsLoader) StartLoading(reader io.Reader, outputNations chan<- CSVNation) {
	csv := csv.NewReader(reader)
	csv.ReuseRecord = true
	csv.FieldsPerRecord = numNationCSVColumns

	for {
		rec, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err == nil && rec != nil {
			nation, err := parseCSVNation(rec)
			if err != nil {
				log.Printf("AiportLoader.StartLoading error: %v %+v\n", err.Error(), nation)
			}

			outputNations <- nation
		}

	}

	close(outputNations)
}

func parseCSVNation(data []string) (CSVNation, error) {
	var result CSVNation

	// CSV structure:
	// NATION,ISO,CURRRENCY,DESCRIPTION
	if len(data) != 4 {
		return result, fmt.Errorf("parseCSVNation: expected num CSV columns 4, got %d", len(data))
	}

	result = NewCSVNation(data[0], data[1], data[2], data[3])
	return result, nil
}
