package lab01

import "fmt"

func Lab01() {
	fmt.Println("Choose exercises from lab01 (1-4): ")
	var ex int
	fmt.Scanf("%d", &ex)

	switch ex {
	case 1:
		Zad01()
	case 2:
		Zad02()
	case 3:
		Zad03()
	case 4:
		Zad04()
	default:
		fmt.Println("You have chosen wrong exercise!")
	}
}
