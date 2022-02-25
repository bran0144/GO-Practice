package main

import (
	"errors"
	"fmt"
)

func divide(divided float64, divsor float64) (float64, error) {
	if divsor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return divided / divsor, nil
}

func main() {
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
