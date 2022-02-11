package chapter8

import (
	"GO-Learning/chapter8/geo"
	"fmt"
)

func main() {
	location := geo.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude =  37.42
	location.Longitude = -122.08
	fmt.Println(location)




}
