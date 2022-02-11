package chapter9

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}
//Need to change the receiver parameter to a pointer type and update the value at the pointer

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	number := Number(4)
	fmt.Println("Original value of a number:", number)
	number.Double()
	fmt.Println("Number after calling double:", number)

	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
