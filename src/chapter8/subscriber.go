package chapter8

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

//Using a struct literal instead
//subscriber := magazine.Subscriber(Name: "John Smith", RateL 4.99, Active: true)
func printInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func defaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.9
	s.Active = true
	return &s
}

func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Signh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Max Ryan")
	printInfo(subscriber2)
	//var s subscriber
	//applyDiscount(&s)
	//fmt.Println(s.rate)
}
