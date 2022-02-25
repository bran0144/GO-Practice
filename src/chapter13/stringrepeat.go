package chapter13

import (
	"fmt"
	"time"
)

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Println(s)
	}
}

func main() {
	go repeat("y")
	go repeat("x")
	time.Sleep(time.Second)
}
