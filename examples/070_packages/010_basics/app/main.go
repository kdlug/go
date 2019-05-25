// https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
package main

// should be run from console
// fo run *.go

// Go first searches for package directory inside GOROOT/src directory and if it doesn’t find the package, then it looks for GOPATH/src
import (
	"course/070_packages/010_basics/greet"
	"course/070_packages/010_basics/greet/de"
	"fmt"
)

var integers [10]int

// Like main function, init function is called by Go when a package is initialized.
// main job of init function is to initialize global variables that can not be initialized in global context.
// For example, initialization of an array.
func init() {
	fmt.Println("app/main.go ===> init() [1])")

	for i := 0; i < 10; i++ {
		integers[i] = i
	}
}

func init() {
	fmt.Println("app/main.go ===> init() [2])")
}

func main() {
	fmt.Println("app/main.go ===> main()")

	fmt.Println(version)
	printVersion() // from version.go file in main package

	fmt.Println(greet.Morning)
	fmt.Println(de.Morning)
}
