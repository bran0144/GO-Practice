//Game that takes input from user to guess a random number between 1-100. User gets 10 guesses
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	//	generate random number (1-100) and store as target number
	target := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guessCount := 0; guessCount < 10; guessCount++{
		fmt.Println("You have", 10 - guessCount, "guesses left.")
		fmt.Print("Enter a guess: ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if guessInt < target {
			fmt.Println("Oops. Your guess was LOW.")
		}
		if guessInt > target {
			fmt.Println("Oops. Your guess was HIGH.")
		}
		if guessInt == target {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess the number. It was:", target)
}
}