// Created by: Mariam Kasim
// Created on: March 2023
//
// This program finds the area of a trapezoid

package main

import "fmt"

func main() {
	// This function finds the area of a trapezoid
	var baseA int
	var baseB int
	var height int
	var area int

	// input
	fmt.Println("This program finds the area of a trapezoid.")
	fmt.Println()
	fmt.Print("Enter the baseA (mm): ")
	fmt.Scanln(&baseA)
	fmt.Print("Enter the baseB (mm): ")
	fmt.Scanln(&baseB)
	fmt.Print("Enter the height (mm): ")
	fmt.Scanln(&height)

	// process
	area = (baseA + baseB) / 2 * height

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "mm².")
	fmt.Println("\nDone.")
}
