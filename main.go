package main

import (
	"fmt"
	"shapes/hollowpyramid"
	"shapes/isoscelespyramid"
	"shapes/isoscelestriangle"
	"shapes/rectangle"
	"shapes/triangle"
)

func main() {
	var length, width int

	fmt.Print("введите длинну стороны прямоугольника или высоты треугольника: ")
	fmt.Scan(&length)
	fmt.Print("введите ширину стороны прямоугольника: ")
	fmt.Scan(&width)

	rectangle.PrintRectangle(width, length)
	fmt.Println("-----------------------------------------------------")
	triangle.PrintTriangle(length)
	fmt.Println("-----------------------------------------------------")
	isoscelestriangle.PrintIsoscelesTriangle(length)
	fmt.Println("-----------------------------------------------------")
	isoscelespyramid.PrintIsoscelesPyramid(length)
	fmt.Println("-----------------------------------------------------")
	hollowpyramid.PrintHollowPyramid(length)

}
