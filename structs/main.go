package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Ways to declare structs
	// info := contactInfo{email: "pit.trak@gmail.com", zip: 17456}
	petros := person{firstName: "Petros", lastName: "Trakadas", contactInfo: contactInfo{email: "pit.trak@gmail.com", zip: 17456}}
	petros.updateName("Pit")
	petros.print()
	// var me person
	// me.firstName = "Petros"
	// me.lastName = "Trakadas"
	// fmt.Println(me)
	// fmt.Printf("%+v", me)
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
