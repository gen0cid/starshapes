package main

import (
	"fmt"
	"shapes/rectangle"
	"shapes/triangle"
)

func main() {
	var length, width int

	fmt.Print("введите длинну: ")
	fmt.Scan(&length)
	fmt.Print("введите ширину: ")
	fmt.Scan(&width)

	rectangle.PrintRectangle(width, length)
	triangle.PrintTriangle(length)

}
