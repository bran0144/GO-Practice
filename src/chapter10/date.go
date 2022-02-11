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

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(0)
	if err != nil {
		log.Fatal(err)
	}

	err1 := date.SetMonth(5)
	if err1 != nil {
		log.Fatal(err1)
	}
	date.SetDay(27)
	err2 := date.SetDay(5)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(date.Year)
	fmt.Println(date.Month)
	fmt.Println(date.Day)
}
