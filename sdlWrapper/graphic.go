package sdlwrapper

import "github.com/veandco/go-sdl2/sdl"

type Graphic struct {
	backgroundColor     *sdl.Color
	title               string
	width               int32
	height              int32
	renderer            *sdl.Renderer
	window              *sdl.Window
	running             bool
	currentColor        *sdl.Color
	timerFPS, lastFrame uint64
	setFPS              uint64
	keystates           []uint8
	mouse               sdl.Point
}

func NewGraphic(title string, width, height int32) (gr *Graphic) {
	gr = &Graphic{title: title, width: width, height: height, setFPS: 60}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, gr.width, gr.height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	gr.window = window

	gr.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	gr.running = true

	return
}

func (gr *Graphic) Finish() {
	gr.running = false
	gr.window.Destroy()
	gr.renderer.Destroy()
	sdl.Quit()
}
