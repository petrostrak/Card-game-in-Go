package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// // Ways to declare structs
	// petros := person{firstName: "Petros", lastName: "Trakadas"}
	// fmt.Println(petros)
	var me person
	me.firstName = "Petros"
	me.lastName = "Trakadas"
	fmt.Println(me)
	fmt.Printf("%+v", me)
}
