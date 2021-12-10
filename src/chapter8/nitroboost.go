package chapter8

import "fmt"

type car2 struct {
	name string
	topSpeed float64
}

func nitroBoost(c *car2) {
	c.topSpeed += 50
}

func main() {
	var mustang car2
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}
