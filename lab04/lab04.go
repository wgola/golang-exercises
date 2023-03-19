package lab04

import "fmt"

func Lab04() {
	fmt.Println("Choose exercise from lab04 (1-7): ")
	var ex int
	fmt.Scanf("%v", &ex)

	switch ex {
	case 1:
		zad01()
	case 2:
		zad02()
	case 3:
		zad03()
	case 4:
		zad04()
	case 5:
		zad05()
	case 6:
		zad06()
	case 7:
		zad07()
	default:
		fmt.Println("You have chosen wrong exercise!")
	}
}
