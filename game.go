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
	guess_count := 0

	reader := bufio.NewReader(os.Stdin)
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
	if guessInt < target{
		guess_count =+1
		fmt.Println("Oops. Your guess was LOW.")
	}
	if guessInt > target {
		guess_count =+1
		fmt.Println("Oops. Your guess was HIGH.")
	}
	if guessInt == target{
		fmt.Println("Good job! You guessed it!")
	}

	}
}