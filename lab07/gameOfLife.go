package lab07

import (
	"math/rand"

	sdlwrapper "golang-exercises/sdlWrapper"
)

type gameOfLife struct {
	life [][]int
}

func NewGameOfLife(width, height int) (gof *gameOfLife) {
	life := [][]int{}
	for i := 0; i < height; i++ {
		tmpRow := []int{}
		for j := 0; j < width; j++ {
			tmpRow = append(tmpRow, 0)
		}
		life = append(life, tmpRow)
	}
	gof = &gameOfLife{life}

	return
}

func (gof *gameOfLife) GetCopyOfBoard() (life [][]int) {
	life = make([][]int, len(gof.life))
	for i := range life {
		life[i] = make([]int, len(gof.life[i]))
		copy(life[i], gof.life[i])
	}

	return
}

func (gof *gameOfLife) RandomizeLife(probability float64) {
	for i := 0; i < len(gof.life[0]); i++ {
		for j := 0; j < len(gof.life); j++ {
			prob := rand.Float64()
			if prob <= probability {
				gof.life[j][i] = 1
			}
		}
	}
}

func (gof *gameOfLife) SimulateTurn(gr *sdlwrapper.Graphic, lifeSize int) {
	boardCopy := gof.GetCopyOfBoard()

	for i := 0; i < len(gof.life[0]); i++ {
		for j := 0; j < len(gof.life); j++ {
			if neighbours := gof.getNeighbours(i, j); gof.life[j][i] == 1 && (neighbours < 2 || neighbours > 3) {
				boardCopy[j][i] = 0
				gr.SetRendererColor(0, 0, 0, 255)
				gr.DrawSquare(int32(i*lifeSize), int32(j*lifeSize), int32(lifeSize))
			} else if gof.life[j][i] == 0 && neighbours == 3 {
				boardCopy[j][i] = 1
				gr.SetRendererColor(255, 255, 0, 255)
				gr.DrawSquare(int32(i*lifeSize), int32(j*lifeSize), int32(lifeSize))
			}
		}
	}

	gof.life = boardCopy
}

func (gof *gameOfLife) getNeighbours(i, j int) (neighbours int) {
	if i-1 >= 0 && gof.life[j][i-1] == 1 {
		neighbours += 1
	}
	if i+1 < len(gof.life[0]) && gof.life[j][i+1] == 1 {
		neighbours += 1
	}
	if j-1 >= 0 && gof.life[j-1][i] == 1 {
		neighbours += 1
	}
	if j+1 < len(gof.life) && gof.life[j+1][i] == 1 {
		neighbours += 1
	}
	if i-1 >= 0 && j-1 >= 0 && gof.life[j-1][i-1] == 1 {
		neighbours += 1
	}
	if i-1 >= 0 && j+1 < len(gof.life) && gof.life[j+1][i-1] == 1 {
		neighbours += 1
	}
	if i+1 < len(gof.life[0]) && j+1 < len(gof.life) && gof.life[j+1][i+1] == 1 {
		neighbours += 1
	}
	if i+1 < len(gof.life[0]) && j-1 >= 0 && gof.life[j-1][i+1] == 1 {
		neighbours += 1
	}

	return
}

func (gof *gameOfLife) DrawBoard(gr *sdlwrapper.Graphic, lifeSize int) {
	for i := 0; i < len(gof.life[0]); i++ {
		for j := 0; j < len(gof.life); j++ {
			if gof.life[i][j] == 1 {
				gr.SetRendererColor(255, 255, 0, 255)
			} else {
				gr.SetRendererColor(0, 0, 0, 255)
			}
			gr.DrawSquare(int32(i*lifeSize), int32(j*lifeSize), int32(lifeSize))
		}
	}
}
