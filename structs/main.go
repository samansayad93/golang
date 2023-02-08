package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	alex.firstName = "John"
	alex.lastName = "Anderson"
	fmt.Println(alex)
}
