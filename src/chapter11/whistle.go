package chapter11

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}
func play (n NoiseMaker) {
	n.MakeSound()
}
type NoiseMaker interface {
	MakeSound()
}
type Robot string
func (r Robot) MakeSound() {
	fmt.Println("Beep boop")
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}

