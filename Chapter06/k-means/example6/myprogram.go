package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	//read data
	irisMatrix := [][]string{}
	iris, err := os.Open("iris.csv")
	if err != nil {
		panic(err)
	}
	defer iris.e()

	reader := csv.NewReader(iris)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		irisMatrix = append(irisMatrix, record)

	}

	X := [][]float64{}
	Y := []string{}
	for _, data := range irisMatrix {
		//convert str-slice into float-slice
		temp := []float64{}
		for _, i := range data[:4] {
			parsedValue, err := strconv.ParseFloat(i, 64)
			if err != nil {
				panic(err)
			}
			temp = append(temp, parsedValue)
		}
		//to explaining variables
		X = append(X, temp)
		//to explained variable
		Y = append(Y, data[4])
	}
	km := Kmeans{}
	fmt.Printf("predict:%v\n", km.fit(X, 3))
	fmt.Printf("teacher:%v\n", Y)
}

type Kmeans struct {
	data            [][]float64
	labels          []int
	representatives [][]float64
}

func Transpose(source [][]float64) [][]float64 {
	var dest [][]float64
	for i := 0; i < len(source[0]); i++ {
		var temp []float64
		for j := 0; j < len(source); j++ {
			temp = append(temp, 0.0)
		}
		dest = append(dest, temp)
	}

	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[0]); j++ {
			dest[j][i] = source[i][j]
		}
	}
	return dest
}

//calculate the distance between two points.Euclidean
func Dist(source, dest []float64) float64 {
	var dist float64
	for i, _ := range source {
		dist += math.Pow(source[i]-dest[i], 2)
	}
	return math.Sqrt(dist)
}

//return index at the point whose value is the smallest on the slice
func ArgMin(target []float64) int {
	var (
		index int
		base  float64
	)
	for i, d := range target {
		if i == 0 {
			index = i
			base = d
		} else {
			if d < base {
				index = i
				base = d
			}
		}

	}
	return index
}

func (km *Kmeans) fit(X [][]float64, k int) []int {
	//set random number seeds
	rand.Seed(time.Now().UnixNano())
	//store data into structure
	km.data = X

	//initialize representative vectors
	//to define the random number range for initialization of representative point, check the max and minimum values of each explaining variables
	transposedData := Transpose(km.data)
	var minMax [][]float64
	for _, d := range transposedData {
		var (
			min float64
			max float64
		)
		for _, n := range d {
			min = math.Min(min, n)
			max = math.Max(max, n)
		}
		minMax = append(minMax, []float64{min, max})
	}
	//set initital values of representative points
	for i := 0; i < k; i++ {
		km.representatives = append(km.representatives, make([]float64, len(minMax)))
	}
	for i := 0; i < k; i++ {
		for j := 0; j < len(minMax); j++ {
			km.representatives[i][j] = rand.Float64()*(minMax[j][1]-minMax[j][0]) + minMax[j][0]
		}
	}
	//initialize label
	//calclate distance between each data and representative point and give label
	for _, d := range km.data {
		var distance []float64
		for _, r := range km.representatives {
			distance = append(distance, Dist(d, r))
		}
		km.labels = append(km.labels, ArgMin(distance))
	}
	for {
		//update representative point
		//set the centroid of the data which belong to the representative point as updated representative point
		//index i means the label
		var tempRepresentatives [][]float64
		for i, _ := range km.representatives {
			var grouped [][]float64
			for j, d := range km.data {
				if km.labels[j] == i {
					grouped = append(grouped, d)
				}
			}
			if len(grouped) != 0 {

				transposedGroup := Transpose(grouped)
				updated := []float64{}
				for _, vectors := range transposedGroup {

					value := 0.0
					for _, v := range vectors {
						value += v
					}
					//store mean of each explaining variables
					updated = append(updated, value/float64(len(vectors)))
				}
				tempRepresentatives = append(tempRepresentatives, updated)
			}
		}
		km.representatives = tempRepresentatives

		//update labels
		tempLabel := []int{}
		for _, d := range km.data {
			var distance []float64
			for _, r := range km.representatives {
				distance = append(distance, Dist(d, r))
			}
			tempLabel = append(tempLabel, ArgMin(distance))
		}
		if reflect.DeepEqual(km.labels, tempLabel) {
			break
		} else {
			km.labels = tempLabel
		}
	}
	return km.labels
}
