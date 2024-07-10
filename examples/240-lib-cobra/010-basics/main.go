// go run main.go --help
// A simple hello world command
//
// Usage:
// hello [flags]
//
// Flags:
// -h, --help          help for hello
// -n, --name string   Name to greet (default "World")

// Example
// go run main.go -n Joe
// Hello, Joe!

// go run main.go Kate
// Hello, World!

package main

import "fmt"

import (
	"github.com/spf13/cobra"
)

// Cobra offers various validation functions for arguments, such as:
// NoArgs - the command does not accept any arguments.
// ExactArgs(n) - the command requires exactly n arguments.
// MinimumNArgs(n) - the command requires at least n arguments.
// MaximumNArgs(n) - the command accepts at most n arguments.
// RangeArgs(min, max) - the command requires between min and max arguments.
func main() {
	var name string

	cmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple hello world command",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	// define string flag
	// StringVarP() - add flag of type string.
	// "P" at the end of the function means, that it supports long and short form of flag (--name or -n)
	// &name - pointer to variable to which the value of the flag will be assigned
	// World - default value
	// Name to greet" - description
	cmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
