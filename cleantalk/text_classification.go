package src

import (

    "fmt"

    "github.com/cdipaolo/goml/base"

    "github.com/cdipaolo/goml/linear"

    "github.com/cdipaolo/goml/text"

    "os"

)

func classifyText(newText string) string {

    modelFile, _ := os.Open("models/model.bin")

    defer modelFile.Close()

    logReg := &linear.Logistic{}

    logReg.Deserialize(modelFile)

    wordDict := text.NewDictionary()

    wordDict.Fit(1000)

    vectorizer := text.NewTfidfVectorizer(wordDict)

    newFeature := vectorizer.Transform(newText).Raw()

    predictedLabel := logReg.Predict([][]float64{newFeature})[0]

    if predictedLabel == 1.0 {

        return "rude"

    } else {

        return "not rude"

    }

}
