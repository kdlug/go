package main

import "errors"

func main() {

}

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		//return "", errors.New("could not find the word you were looking for")
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	// An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type.
	// A gotcha that reference types introduce is that maps can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.
	d[word] = definition
}
