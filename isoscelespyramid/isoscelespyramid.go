package isoscelespyramid

import "fmt"

func PrintIsoscelesPyramid(length int) {
	var widthO = (length * 2) - 1

	for i := 0; i <= length-1; i++ {

		for j := widthO; j > 0; j-- {

			if (j < length-i) || (j > length+i) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
