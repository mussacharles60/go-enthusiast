package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var dadJokes = []string{
	"Why don't eggs tell jokes? They might crack up!",
	"I'm reading a book about anti-gravity. It's impossible to put down!",
	"I used to be a baker, but I couldn't make enough dough.",
	"Why don't skeletons fight each other? They don't have the guts!",
	"What do you call a fake noodle? An impasta!",
}

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Dad Joke Generator!")
	fmt.Println("Enter 'q' to quit.")

	for {
		fmt.Print("Enter any input: ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the user wants to quit
		if input == "q" {
			fmt.Println("Goodbye!")
			break
		}

		// Generate a random dad joke
		joke := getRandomDadJoke()
		fmt.Println("Dad Joke:", joke)
	}
}

func getRandomDadJoke() string {
	index := rand.Intn(len(dadJokes))
	return dadJokes[index]
}
