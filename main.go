// Guess The Number.
// In this game, you guess the number chosen by the app.

package main

import (
	"fmt"
	"math/rand"
)

func setRandomNumber(maxRange int) int {
	fmt.Println("Guess The Number! Choose a number between 0 and", maxRange, "and press enter.")
	return rand.Intn(maxRange)
}

func main() {
	var input int
	maxRange := 5

	randomNumber := setRandomNumber(maxRange)

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
