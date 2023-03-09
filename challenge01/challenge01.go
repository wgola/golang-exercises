package challenge01

import (
	"fmt"
	"golang-exercises/challenge01/stage1"
	"golang-exercises/challenge01/stage2"
)

func Challenge01() {
	fmt.Println("Choose stage (1-8): ")
	var st int
	fmt.Scanf("%v", &st)

	switch st {
	case 1:
		stage1.Stage1()
	case 2:
		stage2.Stage2()
	default:
		fmt.Println("You chose wrong stage!")
	}
}
