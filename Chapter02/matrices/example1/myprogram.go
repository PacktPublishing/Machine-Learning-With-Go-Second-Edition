package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create a flat representation of our matrix.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix.
	a := mat.NewDense(2, 2, data)

	// As a sanity check, output the matrix to standard out.
	fa := mat.Formatted(a, mat.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
