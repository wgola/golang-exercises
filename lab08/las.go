package lab08

import (
	"fmt"
	"math/rand"
)

type Forest struct {
	field [][]int
}

func CreateNewForest(width, height int) (f *Forest) {
	field := [][]int{}
	for i := 0; i < height; i++ {
		tmpRow := []int{}
		for j := 0; j < width; j++ {
			tmpRow = append(tmpRow, 0)
		}
		field = append(field, tmpRow)
	}
	f = &Forest{field}

	return
}

func (f *Forest) GetCopyOfField() (fieldCopy [][]int) {
	fieldCopy = make([][]int, len(f.field))
	for i := range fieldCopy {
		fieldCopy[i] = make([]int, len(f.field[i]))
		copy(fieldCopy[i], f.field[i])
	}

	return
}

func (f *Forest) RandomizeForest(probability float64) {
	for i := 0; i < len(f.field[0]); i++ {
		for j := 0; j < len(f.field); j++ {
			prob := rand.Float64()
			if prob <= probability {
				f.field[j][i] = 1
			}
		}
	}
}

func (f *Forest) RandomizeThunderStrike() {
	y := rand.Intn(len(f.field))
	x := rand.Intn(len(f.field[0]))

	if f.field[y][x] == 1 {
		f.field[y][x] = 2
	}
}

func (f *Forest) SimulateTurn() {
	fieldCopy := f.GetCopyOfField()

	for i := 0; i < len(f.field[0]); i++ {
		for j := 0; j < len(f.field); j++ {
			if f.field[j][i] == 2 {
				fieldCopy = FireNeigbours(fieldCopy, i, j)
			}
		}
	}

	f.field = fieldCopy
}

func (f *Forest) GetNumberOfTrees() (count int) {
	count = 0
	for i := 0; i < len(f.field[0]); i++ {
		for j := 0; j < len(f.field); j++ {
			if f.field[j][i] == 1 {
				count++
			}
		}
	}

	return
}

func (f *Forest) GetNumberOfBurnedTrees() (count int) {
	count = 0
	for i := 0; i < len(f.field[0]); i++ {
		for j := 0; j < len(f.field); j++ {
			if f.field[j][i] == 2 {
				count++
			}
		}
	}

	return
}

func FireNeigbours(fieldCopy [][]int, x, y int) [][]int {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if 0 <= x+j && x+j < len(fieldCopy[0]) && 0 <= y+i && y+i < len(fieldCopy) {
				if fieldCopy[y+i][x+j] == 1 {
					fieldCopy[y+i][x+j] = 2
				}
			}
		}
	}

	return fieldCopy
}

func (f *Forest) DisplayForest() {
	for _, row := range f.field {
		fmt.Println(row)
	}
}
