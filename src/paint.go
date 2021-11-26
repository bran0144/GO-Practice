package main

import "fmt"

func paintNeeded1(width float64, height float64) {
	area := width * height
	fmt.Print("%.2f liters needed \n", area/10.0)
}

//TO use return values:
func paintNeeded2(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}
//To handle errors with this function:
func paintNeeded3(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}
func main() {
		//first example
		paintNeeded1(4.2, 3.0)
		//second exanple
		var amount2, total float64
		amount2 = paintNeeded2(4.2, 3.0)
		fmt.Print("%.2f liters needed \n", amount2)
		total += amount2
		//	third example
		amount3, err := paintNeeded3(4.2, -3.0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%.2f liters needed \n", amount3)
		}

}
