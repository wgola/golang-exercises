package lab04

import (
	"errors"
	"fmt"
)

func zad05() {
	matrix1, matrix2 := generateMatrix(10), generateMatrix(10)

	mulitpied, err := multiplyMatrix(matrix1, matrix2)

	if err != nil {
		fmt.Printf("Error: %s!\n", err)

		return
	}

	fmt.Println("Matrix 1: ")
	printMatrix(matrix1)
	fmt.Println("Matrix 2: ")
	printMatrix(matrix2)
	fmt.Println("Multiplication: ")
	printMatrix(mulitpied)
}

func multiplyMatrix(matrix1, matrix2 [][]int) (multiplied [][]int, err error) {
	if len(matrix1[0]) != len(matrix2) {
		err = errors.New("matrixes cannot be multiplied")

		return
	}

	for i := 0; i < len(matrix1); i++ {
		tmpRow := []int{}
		for j := 0; j < len(matrix2[0]); j++ {
			tmpSum := 0
			for k := 0; k < len(matrix1); k++ {
				tmpSum += matrix1[i][k] * matrix2[k][j]
			}
			tmpRow = append(tmpRow, tmpSum)
		}
		multiplied = append(multiplied, tmpRow)
	}

	return
}
