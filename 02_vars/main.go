package main

import "fmt"

func main() {
	// name := "Hari"
	var age int32 = 43
	const isCool = true
	size := 1.3
	name, email := "Hari", "hk.sudda@gmail.com"
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
