package sdlwrapper

import "github.com/veandco/go-sdl2/sdl"

func (gr *Graphic) SetCurrentColor(r, g, b, a uint8) {
	gr.currentColor = &sdl.Color{R: r, G: g, B: b, A: a}
}

func (gr *Graphic) SetBackgroundColor(r, g, b, a uint8) {
	gr.backgroundColor = &sdl.Color{R: r, G: g, B: b, A: a}
}

func (gr *Graphic) BeginRender() {
	gr.renderer.SetDrawColor(gr.backgroundColor.R, gr.backgroundColor.G, gr.backgroundColor.B, gr.backgroundColor.A)
	gr.renderer.Clear()
	gr.timerFPS = sdl.GetTicks64() - gr.lastFrame
	if gr.timerFPS < (1000 / gr.setFPS) {
		sdl.Delay(uint32((1000 / gr.setFPS) - gr.timerFPS))
	}
	gr.renderer.SetDrawColor(gr.currentColor.R, gr.currentColor.G, gr.currentColor.B, gr.currentColor.A)
}

func (gr *Graphic) EndRender() {
	gr.renderer.Present()
}

func (gr *Graphic) Input() {
	gr.keystates = sdl.GetKeyboardState()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			gr.running = false
		}
	}
	gr.mouse.X, gr.mouse.Y, _ = sdl.GetMouseState()
}

func (gr *Graphic) GetKeystates() []uint8 {
	return gr.keystates
}

func (gr *Graphic) GetRenderer() *sdl.Renderer {
	return gr.renderer
}

func (gr *Graphic) GetRunning() bool {
	return gr.running
}

func (gr *Graphic) SetRunning(running bool) {
	gr.running = running
}

func (gr *Graphic) SetRendererColor(r, g, b, a uint8) {
	gr.renderer.SetDrawColor(r, g, b, a)
}

func (gr *Graphic) DrawSquare(x, y, size int32) {
	gr.renderer.FillRect(&sdl.Rect{X: x, Y: y, W: size, H: size})
}
