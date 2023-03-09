package lab03

import "fmt"

func Lab03() {
	fmt.Println("Choose exercises from lab03 (1-3): ")
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
