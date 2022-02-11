package chapter8

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Using a struct literal instead
//subscriber := magazine.Subscriber(Name: "John Smith", RateL 4.99, Active: true)

type Employee struct{
	Name string
	Salary float64
	Address
//	Can also use anonymous fields and just do Address (no variable)
}

type Address struct {
	Street   string
	City     string
	State    string
	PostCode string
}

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
	address := Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostCode: "68111"}
	subscriber1 := defaultSubscriber("Aman Signh")
	//or you can do this: subscriber1.HomeAddress.City = "Omaha" to set values
	subscriber1.Address = address
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Max Ryan")
	printInfo(subscriber2)
	//var s subscriber
	//applyDiscount(&s)
	//fmt.Println(s.rate)

	var employee Employee
	employee.Name = "Joy Barr"
	employee.Salary = 60000
	employee.Street = "456 Elm St"
	employee.City = "Minneapolis"
	employee.State = "MN"
	employee.PostCode = "55455"



}
