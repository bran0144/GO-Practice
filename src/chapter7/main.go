package chapter7

import (
	"fmt"
	"log"
)

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatalln(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	//var names []string
	//var counts []int
	//for _, line := range lines {
	//	matched := false
	//	for i, name := range names {
	//		if name == line {
	//			counts [i]++
	//			matched = true
	//		}
	//	}
	//	if matched == false {
	//		names = append(names, line)
	//		counts = append(counts, 1)
	//	}
	//}
	//for i, name := range names{
	//	fmt.Printf("%s: %d\n", name, counts[i])
	//}
}
