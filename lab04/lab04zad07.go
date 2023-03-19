package lab04

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/term"
)

func zad07() {
	columns, lines, _ := term.GetSize(0)
	fmt.Print("\033[H\033[2J\033[?25l")

	life := [][]int{}
	for i := 0; i < columns; i++ {
		tmpRow := []int{}
		for j := 0; j < lines; j++ {
			tmpRow = append(tmpRow, 0)
		}
		life = append(life, tmpRow)
	}

	for i := 0; i < columns; i++ {
		for j := 0; j < lines; j++ {
			fmt.Printf("\033[%v;%vH ", i+1, j+1)
		}
	}

	randomizeLife(life)
	for {
		killCells(life)
		createCells(life)
		time.Sleep(1 * time.Second)
	}
}

func randomizeLife(life [][]int) {
	for i := 0; i < len(life[0]); i++ {
		for j := 0; j < len(life); j++ {
			prob := rand.Float64()
			if prob <= 0.35 {
				life[j][i] = 1
				fmt.Printf("\033[%v;%vHO", i+1, j+1)
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
					fmt.Printf("\033[%v;%vH ", i+1, j+1)
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
					fmt.Printf("\033[%v;%vHO", i+1, j+1)
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
