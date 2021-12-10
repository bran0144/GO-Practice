package chapter8

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo2(s student) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Grade: %0.1f\n", s.grade)
}
func main() {
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo2(s)
}
