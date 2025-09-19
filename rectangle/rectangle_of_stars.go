package rectangle

import "fmt"

func PrintRectangle(width, length int) {

	for j := 1; j <= width; j++ {
		for i := 1; i <= length; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
