package gadget

import "fmt"

type Appliance interface {
	TurnOn()
}
type Fan string
func (f Fan) TurnOn() {
	fmt.Println("spinning")
}
type CoffeePot string
func (c CoffeePot) TurnOn() {
	fmt.Println("Powering Up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
}
