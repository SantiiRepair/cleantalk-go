package main

import "fmt"

func main() {

    preprocessData()

    trainModel()

    evaluateModel()

    newText := "This is a rude message."

    predictedLabel := classifyText(newText)

    fmt.Printf("Text: %s\n", newText)

    fmt.Printf("Prediction: %s\n", predictedLabel)

}


