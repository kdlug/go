package main

/*
go run *.go --help

Usage
  -bool
        Prints bool
  -number int
        Prints number
  -string string
        Prints string
  -svar string
        Prints string variable

go run *.go --bool=true -number=7 -string="Hey Joe" -svar="test"
go run *.go --bool=true --number=7 --string="Hey Joe" --svar="test"
*/

/*

As an example, here’s how you parse a string:

var description = flag.String("description", "default value", "the description of the flag")

flag provides many functions to parse different flag types:

func Bool(name string, value bool, usage string) *bool
func BoolVar(p *bool, name string, value bool, usage string)
func Duration(name string, value time.Duration, usage string) *time.Duration
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func Float64(name string, value float64, usage string) *float64
func Float64Var(p *float64, name string, value float64, usage string)
func Int(name string, value int, usage string) *int
func Int64(name string, value int64, usage string) *int64
func Int64Var(p *int64, name string, value int64, usage string)
func IntVar(p *int, name string, value int, usage string)
func String(name string, value string, usage string) *string
func StringVar(p *string, name string, value string, usage string)
func Uint(name string, value uint, usage string) *uint
func Uint64(name string, value uint64, usage string) *uint64
func Uint64Var(p *uint64, name string, value uint64, usage string)
func UintVar(p *uint, name string, value uint, usage string)
*/
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Basic flag declarations are available for string, integer, and boolean options
	// flag.Int, flag.Bool, flag.String functions return a string pointer
	n := flag.Int("number", 0, "Prints number (Required)") // flag name, default option, flag description
	s := flag.String("string", "", "Prints string")
	b := flag.Bool("bool", false, "Prints bool")

	// It's also possible to use existing variable
	var svar string
	// we need to pass in a pointer to the flag declaration function
	flag.StringVar(&svar, "svar", "", "Prints string variable")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// check required flags
	if *n == 0 {
		fmt.Println("number flag is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// we can check types of variable - n,s,b are pointers
	fmt.Printf("n(%T), s(%T), b(%T), svar (%T)\n", n, s, b, svar)

	// we need to dereference the pointers
	fmt.Println(*n)
	fmt.Println(*s)
	fmt.Println(*b)

	fmt.Println(svar)
}
