package lab07

import (
	"flag"
	"time"

	sdlwrapper "golang-exercises/sdlWrapper"

	"github.com/veandco/go-sdl2/sdl"
)

func Lab07zad01() {
	x := flag.Int("X", 800, "Width of windows")
	y := flag.Int("Y", 800, "Height of window")
	size := flag.Int("Z", 4, "Size of one life")
	flag.Parse()

	gr := sdlwrapper.NewGraphic("Game of Life", int32(*x), int32(*y))
	defer gr.Finish()

	gr.SetBackgroundColor(0, 0, 0, 255)
	gr.SetCurrentColor(255, 255, 0, 255)

	life := NewGameOfLife(int(*x / *size), int(*y / *size))
	life.RandomizeLife(0.40)

	for gr.GetRunning() {
		gr.BeginRender()
		life.SimulateTurn(gr, *size)
		time.Sleep(500 * time.Millisecond)
		gr.EndRender()
		gr.Input()
		if gr.GetKeystates()[sdl.SCANCODE_ESCAPE] != 0 {
			gr.SetRunning(false)
		}
	}
}
