package stage3

import (
	"fmt"
)

func Stage3() {
	fmt.Println("Hello! Now you will guess a number which I have randomly picked!")

	for {
		ifStop := game()

		if ifStop {
			break
		} else {
			fmt.Println("Do you want to play again? (Y/n): ")
			var answer string
			fmt.Scanf("%v", &answer)

			if answer == "n" {
				break
			}
		}
	}

	fmt.Println("Thanks for playing!")
}
