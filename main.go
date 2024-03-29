package main

import (
	"fmt"
	"golang-exercises/challenge01"
	"golang-exercises/challenge02"
	"golang-exercises/lab01"
	"golang-exercises/lab03"
	"golang-exercises/lab04"
	"golang-exercises/lab06"
	"golang-exercises/lab08"
)

func main() {
	fmt.Println("Choose lab (1-7) or challenge (100-101): ")
	var lab int
	fmt.Scanf("%d", &lab)

	switch lab {
	case 1:
		lab01.Lab01()
	case 3:
		lab03.Lab03()
	case 4:
		lab04.Lab04()
	case 6:
		lab06.Lab06()
	case 8:
		lab08.LasMain()
	case 100:
		challenge01.Challenge01()
	case 101:
		challenge02.Challenge02()
	default:
		fmt.Println("You have chosen wrong lab!")
	}
}
