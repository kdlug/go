// Why doesn’t this code compile?
// cannot assign to s[0]
package main

import "fmt"

func main() {
	s := "hello"
	s[0] = 'H'
	fmt.Println(s)
}
