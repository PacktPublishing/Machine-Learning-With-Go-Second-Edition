package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the labeled observations and predictions.
	f, err := os.Open("labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// observed and predicted will hold the parsed observed and predicted values
	// form the labeled data file.
	var observed []int
	var predicted []int

	// line will track row numbers for logging.
	line := 1

	// Read in the records looking for unexpected types in the columns.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed and predicted values.
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// classes contains the three possible classes in the labeled data.
	classes := []int{0, 1, 2}

	// Loop over each class.
	for _, class := range classes {

		// These variables will hold our count of true positives and
		// our count of false positives.
		var truePos int
		var falsePos int
		var falseNeg int

		// Accumulate the true positive and false positive counts.
		for idx, oVal := range observed {

			switch oVal {

			// If the observed value is the relevant class, we should check to
			// see if we predicted that class.
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}

				falseNeg++

			// If the observed value is a different class, we should check to see
			// if we predicted a false positive.
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		// Calculate the precision.
		precision := float64(truePos) / float64(truePos+falsePos)

		// Calculate the recall.
		recall := float64(truePos) / float64(truePos+falseNeg)

		// Output the precision value to standard out.
		fmt.Printf("\nPrecision (class %d) = %0.2f", class, precision)
		fmt.Printf("\nRecall (class %d) = %0.2f\n\n", class, recall)
	}
}
