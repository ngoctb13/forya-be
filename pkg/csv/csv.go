package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type RecordProcessor func(record []string, line int) error

func ReadCSV(file io.Reader, expectedFields int, skipHeader bool, process RecordProcessor) ([]error, int, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = expectedFields

	var (
		total   int
		errList []error
	)

	if skipHeader {
		if _, err := reader.Read(); err != nil {
			return nil, 0, fmt.Errorf("invalid csv header: %w", err)
		}
	}

	line := 1
	for {
		line++
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			errList = append(errList, fmt.Errorf("record %d: %w", line, err))
			continue
		}

		total++

		if err := process(record, line); err != nil {
			errList = append(errList, fmt.Errorf("record %d: %w", line, err))
		}
	}

	return errList, total, nil
}
