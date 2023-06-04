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

func trainModel() {

    file, _ := os.Open("data/preprocessed_data.csv")

    defer file.Close()

    csvReader := csv.NewReader(file)

    csvData, _ := csvReader.ReadAll()

    var features [][]float64

    var labels []float64

    for _, row := range csvData {

        feature := vectorizer.Transform(row[0]).Raw()

        label, _ := strconv.ParseFloat(row[1], 64)

        features = append(features, feature)

        labels = append(labels, label)

    }

    XTrain, XTest, yTrain, yTest := base.ClassificationTrainTestSplit(features, labels, 0.2)

    logReg := linear.NewLogistic(base.BatchGA, 0.01, 1000, 1e-6, 0, nil)

    logReg.Fit(XTrain, yTrain)

    trainScore := logReg.Score(XTrain, yTrain, base.Accuracy)

    testScore := logReg.Score(XTest, yTest, base.Accuracy)

    fmt.Printf("Train accuracy: %0.2f\n", trainScore)

    fmt.Printf("Test accuracy: %0.2f\n", testScore)

    modelFile, _ := os.Create("models/model.bin")

    defer modelFile.Close()

    logReg.Serialize(modelFile)

}
