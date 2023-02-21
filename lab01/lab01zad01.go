package main

import "fmt"


func main() {
	fmt.Println("How old are you?")
	
	var i int
	_, err := fmt.Scanf("%d", &i)
	
    if err != nil {
		fmt.Println("Error - wrong age!")
	}

	fmt.Printf("You have lived %d months and %d days!\n", i * 12, i * 12 * 365)
}