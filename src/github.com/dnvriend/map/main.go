package main

import "fmt"

func main() {
	colors := map[string]string {
		"green": "#00ff00",
		"red": "#ff0000",
		"blue": "#0000ff",
	}
	// its mutable!
	colors["white"] = "#ffffff"

	//fmt.Println(colors)
	// it certainly is mutable...
	//delete(colors, "white")
	//fmt.Println(colors)

	printColors(colors)

	barkPurr()
}

func printColors(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func barkPurr() {
	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	for _, value := range m {
		fmt.Println(value)
	}
}
