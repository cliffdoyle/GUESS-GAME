// Package main implements a simple number guessing game.
// The program generates a random number between 1 and 100 and prompts the user to guess it within 10 attempts.
// It provides feedback on each guess (whether it's too low, too high, or correct) and ends the game when the correct number is guessed or when the user runs out of attempts.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator with the current Unix timestamp.
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	// Generate a random number between 1 and 100 (inclusive).
	target := rand.Intn(100) + 1

	// Print a welcome message to the user.
	fmt.Println("I have chosen a random number (1-100). Can you guess it?")

	// Create a reader to read user input from the standard input.
	reader := bufio.NewReader(os.Stdin)

	// Initialize a flag to track if the correct number has been guessed.
	found := false

	// Allow the user 10 guesses.
	for i := 0; i < 10; i++ {
		// Print the number of remaining guesses.
		fmt.Println("You have", 10-i, "guesses left.")
		fmt.Print("Make a guess: ")

		// Read a line of input from the user.
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Remove any leading or trailing whitespace from the input.
		input = strings.TrimSpace(input)

		// Convert the input string to an integer.
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		// Compare the user's guess to the target number and provide feedback.
		if number < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if number > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			// If the guess is correct, set the found flag to true and congratulate the user.
			found = true
			fmt.Println("Congratulations!! You found my number!")
			break
		}
	}

	// If the user did not guess the number, reveal the target number.
	if !found {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
