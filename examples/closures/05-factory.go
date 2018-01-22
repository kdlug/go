/*
Closures can be useful in factories

https://play.golang.org/p/lUMq1DMi0QT
*/
package main

import (
	"fmt"
)

func programmerFactory(lang string) func(name string) {
	return func(programmerName string) {
		fmt.Println(lang, programmerName)
	}
}

func main() {
	phpDevs := programmerFactory("PHP")
	phpDevs("Kate")
	phpDevs("John")
	phpDevs("Alec")

	goDevs := programmerFactory("GO")
	goDevs("George")
	goDevs("Joe")
}
