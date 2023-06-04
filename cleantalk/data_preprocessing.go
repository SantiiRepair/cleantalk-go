package src

import (

    "encoding/csv"

    "os"

    "strings"

)

func preprocessData() {

    file, _ := os.Open("data/raw_data.csv")

    defer file.Close()

    csvReader := csv.NewReader(file)

    csvData, _ := csvReader.ReadAll()

    preprocessedData := [][]string{}

    for _, row := range csvData {

        preprocessedRow := []string{}

        for _, word := range strings.Fields(row[0]) {

            if !base.StopWordMap[strings.ToLower(word)] {

                preprocessedRow = append(preprocessedRow, strings.ToLower(word))

            }

        }

        preprocessedData = append(preprocessedData, preprocessedRow)

    }

    outputFile, _ := os.Create("data/preprocessed_data.csv")

    defer outputFile.Close()

    csvWriter := csv.NewWriter(outputFile)

    csvWriter.WriteAll(preprocessedData)

}
	
