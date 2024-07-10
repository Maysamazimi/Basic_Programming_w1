package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	var height float64

	// Place to enter radius
	fmt.Print("Enter the radius of the tube: ")
	fmt.Scanln(&radius)

	// Place to enter height
	fmt.Print("Enter the height of the tube: ")
	fmt.Scanln(&height)

	// Formula to Calculate the area of the tube
	area := math.Pi * radius * (radius + height)

	fmt.Printf("Surface area of the tube: %.1f\n", area)
}
