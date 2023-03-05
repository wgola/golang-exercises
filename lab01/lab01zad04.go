package lab01

import (
	"fmt"
	"math/rand"
)

func Zad04() {
	answer := rand.Intn(100) + 1

	fmt.Println("Guess the number (1-100) or type 0 to exit: ")

	var guess int
	fmt.Scanf("%d", &guess)

	var end bool

	for !end {
		if guess == 0 {
			fmt.Println("See you later!")
			end = true
		} else if guess == answer {
			fmt.Println("Congrats! You guessed the answer!")
			end = true
		} else {
			if guess < answer {
				fmt.Println("Your guess is too little!")
			} else {
				fmt.Println("Your guess is too big!")
			}
			fmt.Println("Try again or type 0 to exit")
			fmt.Scanf("%d", &guess)
		}
	}
}
