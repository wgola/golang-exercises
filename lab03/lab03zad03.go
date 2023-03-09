package lab03

import (
	"github.com/sbinet/go-gnuplot"
)

func Zad03() {
	var (
		numbers = [100000]float64{}
		lengths = [100000]float64{}
		fname   = ""
		persist = true
		debug   = false
	)

	for i := 1; i <= 100000; i++ {
		x, y := LongestCycle(i, i)
		numbers[i-1], lengths[i-1] = float64(x), float64(y)
	}

	p, _ := gnuplot.NewPlotter(fname, persist, debug)
	defer p.Close()

	p.SetLabels("numbers", "lengths")
	p.SetStyle("dots")
	p.CheckedCmd("set term png size 1280, 720")
	p.CheckedCmd("set output 'plot.png'")
	p.PlotXY(numbers[:], lengths[:], "Lengths of Collatz conjecture for numbers in range 1 - 100000")
	p.CheckedCmd("q")
}
