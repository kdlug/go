// Why doesn’t these print statements give the same result?
package main

import "fmt"

func main() {
	fmt.Println("H" + "i")
	fmt.Println('H' + 'i')
}
