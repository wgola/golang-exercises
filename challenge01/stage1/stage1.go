package stage1

import (
	"fmt"
	"math/rand"
)

func Stage1() {
	fmt.Println("Hello! Now you will guess a number which I have randomly picked!")
	randomNumber := rand.Intn(100) + 1

	fmt.Println("Enter a number (1-100): ")

	var givenNumber int
	fmt.Scanf("%v", &givenNumber)

	for givenNumber != randomNumber {
		if givenNumber < randomNumber {
			fmt.Println("Your number is too little!")
		} else {
			fmt.Println("Your number is too big!")
		}
		fmt.Println("Try another number (1-100): ")

		fmt.Scanf("%v", &givenNumber)
	}

	fmt.Printf("Congratulations! The number was %v!\n", randomNumber)
}
