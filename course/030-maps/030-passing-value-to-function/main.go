// Map is a reference type so we don't need to pass a variable by refference to change it's value
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff", // comma is mandatory at the end
	}

	changeColor(colors, "green", "#00ff01")

	fmt.Println(colors) // map[red:#ff0000 green:#00ff01 white:#ffffff]
}

// without pointer
func changeColor(c map[string]string, color string, hex string) {
	c[color] = hex
}
