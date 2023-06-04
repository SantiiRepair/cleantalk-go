package src

import (

    "encoding/csv"

    "fmt"

    "github.com/cdipaolo/goml/base"

    "github.com/cdipaolo/goml/linear"

    "github.com/cdipaolo/goml/text"

    "os"

    "strconv"

)

func evaluateModel() {

    file, _ := os.Open("data/train_test_data.csv")

    defer file.Close()

    csvReader := csv.NewReader(file)

    csvData, _ := csvReader.ReadAll()

    modelFile, _ := os.Open("models/model.bin")

    defer modelFile.Close()

    logReg := &linear.Logistic{}

    logReg.Deserialize(modelFile)

    wordDict := text.NewDictionary()

    wordDict.Fit(1000)

    vectorizer := text.NewTfidfVectorizer(wordDict)

    var features [][]float64

    var labels []float64

    for _, row := range csvData {

        feature := vectorizer.Transform(row[0]).Raw()

        label, _ := strconv.ParseFloat(row[1], 64)

        features = append(features, feature)

        labels = append(labels, label)

    }

    score := logReg.Score(features, labels, base.Accuracy)

    fmt.Printf("Test accuracy: %0.2f\n", score)

}
