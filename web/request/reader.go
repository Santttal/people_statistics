package request

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

type PersonRecordsReader interface {
	Read(r io.Reader) ([]PersonRecord, error)
}

type CsvReaderWrapper struct {
}

func (w CsvReaderWrapper) Read(r io.Reader) ([]PersonRecord, error) {
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	personRecords := make([]PersonRecord, len(lines)-1)
	var name string
	var gender string
	var age int

	for i, line := range lines {
		if len(line) != 3 {
			return nil, errors.New("invalid line size")
		}
		if i == 0 {
			// validate headers
			continue
		}
		name = line[0]
		gender = line[1]
		if age, err = strconv.Atoi(line[2]); err != nil {
			return nil, err
		}
		personRecords[i-1] = PersonRecord{name, gender, age}
	}

	return personRecords, nil
}
