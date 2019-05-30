package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)

func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastname
	}
}
func main() {
	fmt.Println("Structs:")
	//person1 := Person{firstName: "Hari", lastName: "Sudda", city: "melbourne", gender: "m", age: 38}
	person1 := Person{"Naga", "bathula", "trug", "f", 36}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Hari")
	fmt.Println(person1.greet())

}
