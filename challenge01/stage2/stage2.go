package stage2

import (
	"fmt"
	"math/rand"
)

func Stage2() {
	fmt.Println("Hello! Now you will guess a number which I have randomly picked!")
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
		fmt.Println("Goodbye!")
	} else {
		fmt.Printf("Congratulations! The number was %v!\n", randomNumber)
	}
}
