package chapter8

import "fmt"

type part2 struct {
	description string
	count int
}

func doublePack( p *part2) {
	p.count +=2
}
func main() {
	var fuses part2
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}
