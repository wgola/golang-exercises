package main

import (
	"fmt"
	"golang-exercises/lab01"
)

func main() {
	fmt.Println("Choose lab (1): ")
	var lab int
	fmt.Scanf("%d", &lab)

	switch lab {
	case 1:
		lab01_exercises()
	default:
		fmt.Println("You have chosen wrong lab!")
	}
}

func lab01_exercises() {
	fmt.Println("Choose exercises from lab01 (1-4): ")
	var ex int
	fmt.Scanf("%d", &ex)

	switch ex {
	case 1:
		lab01.Zad01()
	case 2:
		lab01.Zad02()
	case 3:
		lab01.Zad03()
	case 4:
		lab01.Zad04()
	default:
		fmt.Println("You have chosen wrong exercise!")
	}
}
