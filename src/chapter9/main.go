package chapter9

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64


func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
//can create a ToGallons method for Milliliters too, but need to change syntax for other method too
func (m Milliliters) ToGallons() Gallons{
	return Gallons(m * 0.000264)
}

func (g Gallons)ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {

	soda := Liters(2)
	fmt.Println("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Println("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
	milk := Gallons(2)
	fmt.Println("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	l := Liters(3)
	fmt.Println("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Println("%0.1f milliliters is %0.1f liters\n", ml, ml.ToLiters())
	//carFuel := Gallons(1.2)
	//busFuel := Liters(4.5)
	//carFuel += ToGallons(Liters(40.0))
	//busFuel += ToLiters(Gallons(30.0))
	//fmt.Println("Car fuel: %0.lf gallons\n", carFuel)
	//fmt.Println("Bus fuel: %0.lf liters\n", busFuel)
}
