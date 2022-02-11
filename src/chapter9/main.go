package chapter9

import "fmt"

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)
	fmt.Println(carFuel, busFuel)
	//conversion
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Println(carFuel, busFuel)
}
