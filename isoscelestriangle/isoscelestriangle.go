package isoscelestriangle

import "fmt"

func PrintIsoscelesTriangle(length int) {
	for i := 1; i <= length; i++ {
		for j := length; j > 0; j-- {

			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Println()
	}
}
