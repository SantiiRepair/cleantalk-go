package main

import (

    "github.com/cdipaolo/goml/base"

    "github.com/cdipaolo/goml/linear"

)

func trainModel() {



    data := base.LoadDataFromCSV("data/text_data.csv")



    pipeline := base.NewPipeline(

        base.TFIDF{},

    )



    model := linear.NewL2LogisticRegression(0.01, 100, 1e-6, true)

    pipeline.Add(model)

    pipeline.Fit(data)



    pipeline.Save("models/trained.bin")

}
