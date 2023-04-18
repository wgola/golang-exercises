package lab07

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	bkg                 sdl.Color
	title               string = "SDL2 Window"
	width               int32  = 800
	height              int32  = 800
	renderer            *sdl.Renderer
	window              *sdl.Window
	timerFPS, lastFrame uint64
	setFPS              uint64 = 60
	mouse               sdl.Point
	keystates           = sdl.GetKeyboardState()
	running             bool
)

func setColor(r, g, b, a uint8) sdl.Color {
	var c sdl.Color
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	return c
}

func start() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	running = true
}

func startSet(t string, w int32, h int32) {
	title = t
	width = w
	height = h
	start()
}

func quit() {
	running = false
	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}

func input() {
	keystates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}
	}
	mouse.X, mouse.Y, _ = sdl.GetMouseState()
}

func beginRender() {
	renderer.SetDrawColor(bkg.R, bkg.G, bkg.B, bkg.A)
	renderer.Clear()
	timerFPS = sdl.GetTicks64() - lastFrame
	if timerFPS < (1000 / setFPS) {
		sdl.Delay(uint32((1000 / setFPS) - timerFPS))
	}
	renderer.SetDrawColor(0, 0, 0, 255)
}

func endRender() {
	renderer.Present()
}

func Lab07zad01() {
	startSet("Game of Life", 1200, 1200)
	bkg = setColor(0, 0, 0, 255)
	life := [][]int{}
	for i := 0; i < int(width/10); i++ {
		tmpRow := []int{}
		for j := 0; j < int(height/10); j++ {
			tmpRow = append(tmpRow, 0)
		}
		life = append(life, tmpRow)
	}

	randomizeLife(life)
	for running {
		beginRender()
		killCells(life)
		createCells(life)
		time.Sleep(1 * time.Second)
		endRender()
		input()
		if keystates[sdl.SCANCODE_ESCAPE] != 0 {
			running = false
		}
	}
	quit()
}

func randomizeLife(life [][]int) {
	for i := 0; i < len(life[0]); i++ {
		for j := 0; j < len(life); j++ {
			prob := rand.Float64()
			if prob <= 0.5 {
				life[j][i] = 1
				renderer.SetDrawColor(255, 255, 0, 255)
				renderer.FillRect(&sdl.Rect{X: int32(i * 10), Y: int32(j * 10), W: 10, H: 10})
			}
		}
	}
}

func killCells(life [][]int) {
	for i := 0; i < len(life[0]); i++ {
		for j := 0; j < len(life); j++ {
			if life[j][i] == 1 {
				neighbours := getNeighbours(i, j, life)
				if neighbours < 2 || neighbours > 3 {
					life[j][i] = 0
					renderer.SetDrawColor(0, 0, 0, 255)
					renderer.FillRect(&sdl.Rect{X: int32(i * 10), Y: int32(j * 10), W: 10, H: 10})
				}
			}
		}
	}
}

func createCells(life [][]int) {
	for i := 0; i < len(life[0]); i++ {
		for j := 0; j < len(life); j++ {
			if life[j][i] == 0 {
				neighbours := getNeighbours(i, j, life)
				if neighbours == 2 {
					life[j][i] = 1
					renderer.SetDrawColor(255, 255, 0, 255)
					renderer.FillRect(&sdl.Rect{X: int32(i * 10), Y: int32(j * 10), W: 10, H: 10})
				}
			}
		}
	}
}

func getNeighbours(i, j int, life [][]int) (neighbours int) {
	if i-1 >= 0 && life[j][i-1] == 1 {
		neighbours += 1
	}
	if i+1 < len(life[0]) && life[j][i+1] == 1 {
		neighbours += 1
	}
	if j-1 >= 0 && life[j-1][i] == 1 {
		neighbours += 1
	}
	if j+1 < len(life) && life[j+1][i] == 1 {
		neighbours += 1
	}

	return
}
