package chapter8

import (
	"GO-Learning/chapter8/geo"
	"fmt"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
