package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create a new matrix a.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}
