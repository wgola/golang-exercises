package stage3

import (
	"fmt"
	"math/rand"
)

func game() bool {
	randomNumber := rand.Intn(100) + 1
	fmt.Println("Enter a number (1-100) (or 0 to end program): ")

	var givenNumber int
	fmt.Scanf("%v", &givenNumber)

	for givenNumber != 0 && givenNumber != randomNumber {
		if givenNumber < randomNumber {
			fmt.Println("Your number is too little!")
		} else {
			fmt.Println("Your number is too big!")
		}
		fmt.Println("Try another number (1-100) (or 0 to end program): ")

		fmt.Scanf("%v", &givenNumber)
	}

	if givenNumber == 0 {
		return true
	} else {
		fmt.Printf("Congratulations! The number was %v!\n", randomNumber)
		return false
	}
}
