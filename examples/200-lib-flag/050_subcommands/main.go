/*
Subcommands can be useful when your app needs to perform multiple functions, and each function has a unique set of flags and arguments.
A common example of a CLI tool with subcommands is git:

git **commit** -m "message" git **push**

Subcommands are implemented by creating a flag.FlagSet.
The FlagSet has a subcommand name, error handling behavior, and a set of flags associated with it.
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

var count int
var list string

var commandCountVar string
var commandMetricVar string
var commandSubstringVar string
var commandUniqueVar bool

var listMetricVar string
var listTextVar string
var listUniqueVar bool

// var countCommand *flag.FlagSet
// var listCommand *flag.FlagSet

func init() {
}

func main() {
	// subcommands
	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	// Count subcommand flag pointers
	countCommand.StringVar(&commandCountVar, "text", "", "Text to parse. (Required)")
	countCommand.StringVar(&commandMetricVar, "metric", "chars", "Metric {chars|words|lines|substring} (Required)")
	countCommand.StringVar(&commandSubstringVar, "text", "", "The substring to be counted. Required for --metric=substring")
	countCommand.BoolVar(&commandUniqueVar, "unique", false, "Measure unique values of a metric.")

	listCommand.StringVar(&listTextVar, "text", "", "Text to parse. (Required)")
	listCommand.StringVar(&listMetricVar, "metric", "chars", "Metric <chars|words|lines>. (Required)")
	listCommand.BoolVar(&listUniqueVar, "unique", false, "Measure unique values of a metric.")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}
	// Switch on the subcommand
	// 	// Parse the flags for appropriate FlagSet
	// 	// FlagSet.Parse() requires a set of arguments to parse as input
	// 	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
	case "count":
		countCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
