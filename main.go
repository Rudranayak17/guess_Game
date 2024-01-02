package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess the number between 1 and 10:")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	generateRandom := random.Intn(10) + 1

	attempts := 3

	for i := 1; i <= attempts; i++ {
		fmt.Printf("Attempt %d: Enter a number: ", i)
		var userNumber int
		_, err := fmt.Scan(&userNumber)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if userNumber == generateRandom {
			fmt.Printf("Congratulations, you guessed the correct number %d! You took %d attempts.\n", generateRandom, i)
			return
		} else if userNumber > generateRandom {
			fmt.Printf("Too high! ")
		} else {
			fmt.Printf("Too low! ")
		}

		fmt.Printf("You have %d attempts remaining.\n", attempts-i)
	}

	fmt.Printf("Sorry, you're out of attempts. The correct number was %d.\n", generateRandom)
}
