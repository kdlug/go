package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// func Open(name string) (file *File, err error)
	// os.Open function returns a non-nil error value when it fails to open a file

	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}

	// this code won't be run if fatal occurs
	// do something with the open *File f
	fmt.Println(f)

}
