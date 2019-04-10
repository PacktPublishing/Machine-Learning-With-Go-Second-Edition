package main

import (
	"image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the advertising dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(f)

	// pts will hold the values for plotting
	pts := make(plotter.XYs, irisDF.Nrow())
	yVals := irisDF.Col("sepal_width").Float()

	for i, floatVal := range irisDF.Col("sepal_length").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// Create the plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "sepallength"
	p.Y.Label.Text = "sepalwidth"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Save the plot to a PNG file.
	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
		log.Fatal(err)
	}

}
