package main

func main() {
	/* panic: assignment to entry in nil map

	var cars map[string]string
	cars["honda"] = "CRV"
	*/

	var cars map[string]string
	// You have to initialize the map using the make function (or a map literal) before you can add any elements
	cars = make(map[string]string)
	cars["honda"] = "CRV"

}
