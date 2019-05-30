package main

import "fmt"

func main() {
	fmt.Println("Maps:")
	// Define map
	emails := make(map[string]string)

	// emails := map[string]string{"Bob": "bob@gmail.com"}
	//Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["rob"] = "rob@gmail.com"
	emails["Mike"] = "mkike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
	delete(emails, "Bob")

	fmt.Println(emails)

}
