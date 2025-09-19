package triangle

import "fmt"

func PrintTriangle(v int) {
	for i := 1; i <= v; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
