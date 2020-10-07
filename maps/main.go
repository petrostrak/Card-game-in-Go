package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#008000",
	// 	"blue":  "#0000FF",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	delete(colors, "white")
	fmt.Println(colors)
}
