package sPlot

import (
	"image/color"

	"github.com/yasseldg/simplego/sLog"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func SaveHistogram(dist []float64, title, filePath string) {

	n := len(dist)
	vals := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		vals[i] = dist[i]
	}

	plt := plot.New()
	plt.Title.Text = title
	hist, err := plotter.NewHist(vals, 25) // 25 bins
	if err != nil {
		sLog.Error("Cannot plot: %s ", err)
	}
	hist.FillColor = color.RGBA{R: 255, G: 127, B: 80, A: 255} // coral color
	plt.Add(hist)

	err = plt.Save(400, 200, filePath)
	if err != nil {
		sLog.Panic(err.Error())
	}
}
