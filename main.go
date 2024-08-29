// Guess The Number.
// In this game, you guess the number chosen by the app.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	randomNumber := rand.Intn(5)
	fmt.Println("Guess The Number! Choose a number between 0 and 5 and press enter.")
	for {
		fmt.Scanln(&input)
		if input != randomNumber {
			fmt.Println("Wrong! Try again.")
		} else {
			fmt.Println("You are right!")
			break
		}
	}
}
