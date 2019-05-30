package main

import "fmt"

func main() {
	fmt.Println("Pointers:")
	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// user * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change val with pointer

	*b = 10
	fmt.Println(a)
}
