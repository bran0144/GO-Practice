package chapter10

import "fmt"

type Date struct {
	Year int
	Month int
	Day int
}

func (d Date) SetYear(year int) {
	d.Year = year
}

func main() {
	date := Date{Year: 2019, Month: 5, Day: 27}
	fmt.Println(date)
}
