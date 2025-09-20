package main

import (
	"fmt"
	"shapes/isoscelestriangle"
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
	fmt.Println("-----------------------------------------------------")
	triangle.PrintTriangle(length)
	fmt.Println("-----------------------------------------------------")
	isoscelestriangle.PrintIsoscelesTriangle(length)
}
