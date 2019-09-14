package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

type SingleRecord map[string]string

type Callback func(record SingleRecord, args interface{}) error

func (c *CsvRecord) GetInt(field string) (int, error) {

	var r int
	var err error

	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		return 0, err
	}
	return r, nil
}

func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	} else {
		return ""
	}
}

func LoadCsvCfg(filename string, row int) (*CsvTable, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		return nil, err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < row {
		return nil, err
	}

	colNum := len(records[0])
	recordNum := len(records)

	var allRecords []CsvRecord

	for i := row; i < recordNum; i++ {

		record := &CsvRecord{make(map[string]string)}

		for k := 0; k < colNum; k++ {
			record.Record[records[0][k]] = records[i][k]
		}

		allRecords = append(allRecords, *record)
	}

	var result = &CsvTable{
		filename,
		allRecords,
	}
	return result, nil
}

//DealCsvByRow
func DealCsvByRow(filename string, callback Callback, args interface{}) error {

	file, err := os.Open(filename)

	rows := 0

	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var header []string
	record := make(SingleRecord)

	colNum := 0

	for {

		rows++

		read, err := reader.Read()

		if rows == 1 {
			colNum = len(read)
			header = read
			continue
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		for k := 0; k < colNum; k++ {
			record[header[k]] = read[k]
		}

		_ = callback(record, args)
	}

	return nil

}
