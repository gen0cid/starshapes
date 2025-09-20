package triangle

import "fmt"

func PrintTriangle(length int) {
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
