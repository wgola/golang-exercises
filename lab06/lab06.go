package lab06

import (
	"fmt"
)

func Lab06() {
	fmt.Println("Choose exercise from lab06 (1): ")
	var ex int
	fmt.Scanf("%v", &ex)

	switch ex {
	case 1:
		zad01()
	default:
		fmt.Println("You have chosen wrong exercise!")
	}
}
