package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person { firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v", alex)
}
