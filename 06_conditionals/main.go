package main

import "fmt"

func main() {
	fmt.Println("Conditionals:")
	x := 5
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than  %d\n", x, y)
	} else {
		fmt.Println("Otherway round")
	}

	// if else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	//switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is Not red and blue")
	}
}
