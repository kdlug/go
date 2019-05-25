package main

import "fmt"

func main() {
	// simple anonymous function without name
	func() {
		fmt.Println("Hi there")
	}() // () invoke the function immediately

	for i := 0; i < 5; i++ {
		func() {
			// function has an access to external i variable
			// but it will cause a problem if we call this func asynchrounously
			fmt.Println(i)
		}()

		// so best practice is to passing an args
		func(i int) {
			// now we're using local variable
			fmt.Println(i)
		}(i) // passing argument when calling
	}

	// we can pass a function to a variable
	// var f func() = func() ...
	f := func() {
		fmt.Println("Hey Joe!")
	} // we don't call it here

	// we are calling it here
	f()
}
