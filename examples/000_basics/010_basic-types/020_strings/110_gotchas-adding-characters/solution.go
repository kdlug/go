// Why doesn’t these print statements give the same result?
package main

import "fmt"

func main() {
	fmt.Println("H" + "i")
	// The rune literals 'H' and 'i' are integer values identifying Unicode code points: 'H' is 72 and 'i' is 105.
	fmt.Println('H' + 'i')

	// How to convert rune/int to a string?
	fmt.Println(string('H') + string('i'))
	fmt.Println(string(72) + string(105))

	// or using fmt.Sprintf function
	s := fmt.Sprintf("%c%c, world!", 72, 'i')
	fmt.Println(s) // "Hi, world!"
}
