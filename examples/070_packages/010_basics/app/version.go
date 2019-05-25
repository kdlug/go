package main

import "fmt"

// A scope is a region in code block where a defined variable is accessible. A package scope is a region within a package where a declared variable is accessible from within a package (across all files in the package).
// This region is the top-most block of any file in the package.
var version = "1.0.0"

func printVersion() {
	fmt.Println(version)
}
