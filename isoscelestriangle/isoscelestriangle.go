package isoscelestriangle

import "fmt"

func PrintIsoscelesTriangle(length int) {
	for i := 1; i <= length; i++ {
		for j := 1; j <= length-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
