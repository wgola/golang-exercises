package main

import (
	"fmt"
	"golang-exercises/challenge01"
	"golang-exercises/lab01"
	"golang-exercises/lab03"
)

func main() {
	fmt.Println("Choose lab (1-3) or challenge (100-101): ")
	var lab int
	fmt.Scanf("%d", &lab)

	switch lab {
	case 1:
		lab01.Lab01()
	case 3:
		lab03.Lab03()
	case 100:
		challenge01.Challenge01()
	default:
		fmt.Println("You have chosen wrong lab!")
	}
}
