package chapter12

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk")
	fmt.Println("Nice weather!")
	return nil
}
func main () {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
