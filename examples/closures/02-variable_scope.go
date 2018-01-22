/*
Scope of variables
https://play.golang.org/p/K8SLkQqLAT6
*/
package main

import (
	"fmt"
)

func localVarScope() {
	outsideVar := 100

	for i := 0; i < 1; i++ {
		insideVar := 200
		fmt.Println(outsideVar) // 100
		fmt.Println(insideVar)  // 200
	}

	// fmt.Println(insideVar) // is out of scrope, undefined: insideVar
}

func main() {
	localVarScope()

	// fmt.Println(outsideVar) // is out of scrope, undefined: outsideVar
	// if we use anonymous function, we can access outsideVar
}

