package jsontocsv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func JsonToCSV(jsonFileName string, csvFileName string) error {
	// Read the JSON file
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var ranking []map[string]interface{}
	if err := json.NewDecoder(jsonFile).Decode(&ranking); err != nil {
		return err
	}

	// Create a CSV file
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	var header []string
	for k := range ranking[0] {
		header = append(header, k)
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	// 4. Write the successive rows by iterating through the map array and converting the values to strings
	for _, r := range ranking {
		var csvRow []string
		for _, h := range header {
			v := r[h]
			switch v.(type) {
			case string:
				csvRow = append(csvRow, v.(string))
			case float64:
				csvRow = append(csvRow, fmt.Sprint(v.(float64)))
			case bool:
				csvRow = append(csvRow, fmt.Sprint(v.(bool)))
			default:
				csvRow = append(csvRow, "")
			}
		}
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}

	return nil
}
