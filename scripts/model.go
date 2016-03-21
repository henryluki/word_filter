package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	// "github.com/sjwhitworth/golearn/evaluation"
	// "github.com/sjwhitworth/golearn/linear_models"
)

func main() {
	rawData, err := base.ParseCSVToInstances("../data/pre/training.csv", true)
	if err != nil {
		panic(err)
	}

	fmt.Println(rawData)
	// cls := linear_models.NewLogisticRegression("l1", 0.50, 100)

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	fmt.Println(trainData)
	fmt.Println(testData)
	// cls.Fit(trainData)

	// Calculates the Euclidean distance and returns the most popular label
	// predictions := cls.Predict(testData)
	// fmt.Println(predictions)

	// Prints precision/recall metrics
	// confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	// if err != nil {
	// panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	// }
	// fmt.Println(evaluation.GetSummary(confusionMat))
}
