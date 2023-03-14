package lab04

import "fmt"

func zad01() {
	tab1 := [20]float64{}
	tab2 := [20]float64{}

	for i := 0; i < len(tab1); i++ {
		tab1[i] = 2.0
		tab2[i] = 3.0
	}

	sum := [20]float64{}
	for i := 0; i < len(tab1); i++ {
		sum[i] = tab1[i] + tab2[i]
	}

	fmt.Printf("Array 1: %v\nArray 2: %v\nSum of arrays: %v\n", tab1, tab2, sum)
}
