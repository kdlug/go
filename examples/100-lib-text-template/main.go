package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift string
	Attended   bool
}

var recipients = []Recipient{
	{"Aunt Mildred", "bone china tea set", true},
	{"Uncle John", "moleskin pants", false},
	{"Cousin Rodney", "", false},
}

func main() {
	// load template from a file
	filename := "letter.tmpl"
	tmpl := template.Must(template.ParseFiles(filename))

	// define a buffer writer
	var writer bytes.Buffer

	// Execute the template for each recipient.
	for _, r := range recipients {

		err := tmpl.Execute(&writer, r) // we need to pass a pointer (address) to writer
		if err != nil {
			log.Fatal(err)
		}
	}

	// return buffer a string
	result := writer.String()

	fmt.Println(result)
}
