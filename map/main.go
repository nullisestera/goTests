package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FA8072",
		"green": "#228B22",
		"blue":  "#1E90FF",
	}

	fmt.Println(colors["red"])

	// another way

	colors2 := make(map[string]string)

	colors2["white"] = "#FFFFFF"
	colors2["cyan"] = "#40E0D0"

	fmt.Println(colors2["cyan"])

	// deleting
	delete(colors, "blue")
	fmt.Println(colors)

	printMap(colors)
}

// Iterating over map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("print map:", color+" "+hex)
	}
}
