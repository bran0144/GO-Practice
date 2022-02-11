package chapter10

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

//Setter methods need pointer receivers (don't need to do anything to d
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	date := Date{}
	err := date.SetYear(0)
	if err != nil {
		log.Fatal(err)
	}
	date.SetMonth(5)
	date.SetDay(27)
	fmt.Println(date.Year)
}
