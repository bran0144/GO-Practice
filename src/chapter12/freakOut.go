package chapter12

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}
func freakOut() {
	panic("oh no")
	fmt.Println("I won't be run!")
}
func main() {
	defer calmDown()
	freakOut()
	fmt.Println("Exiting normally")
}
