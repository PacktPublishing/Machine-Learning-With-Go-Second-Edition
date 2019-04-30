package main

import (
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the advertising dataset file.
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(f)

	// Extract the target column.
	yVals := advertDF.Col("Sales").Float()

	// Create a scatter plot for each of the features in the dataset.
	for _, colName := range advertDF.Names() {

		// pts will hold the values for plotting
		pts := make(plotter.XYs, advertDF.Nrow())

		// Fill pts with data.
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// Create the plot.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		// Save the plot to a PNG file.
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
