package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	sepalLength := irisDF.Col("sepal_length").Float()
	sepalWidth := irisDF.Col("sepal_width").Float()
	cor, _ := stats.Correlation(sepalLength, sepalWidth)
	fmt.Println(cor)
}
