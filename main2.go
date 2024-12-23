package main

import (
	"fmt"
	"github.com/Arsencchikkk/BayakhmetArsen-assignment1-Library/Geometria"
)

func main() {
	var choice int

	fmt.Println("Choose the shape to calculate:")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Circle")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var length, width float64
		fmt.Print("Enter the length of the rectangle: ")
		fmt.Scan(&length)
		fmt.Print("Enter the width of the rectangle: ")
		fmt.Scan(&width)

		rectangle := Geometria.Rectangle{Length: length, Width: width}
		fmt.Println("\nRectangle details:")
		Geometria.PrintShapeDetails(rectangle)

	case 2:
		var radius float64
		fmt.Print("Enter the radius of the circle: ")
		fmt.Scan(&radius)

		circle := Geometria.Circle{Radius: radius}
		fmt.Println("\nCircle details:")
		Geometria.PrintShapeDetails(circle)

	default:
		fmt.Println("Invalid choice.")
	}
}
