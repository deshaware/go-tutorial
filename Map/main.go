package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff00000",
	// 	"green": "#4bf745",
	// }

	// other way
	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// colors["10"] = "#0000"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for code := range c {
		fmt.Println("Hex code for", "is ", code)
	}
}
