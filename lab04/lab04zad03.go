package lab04

import "fmt"

func zad03() {
	matrix := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}
