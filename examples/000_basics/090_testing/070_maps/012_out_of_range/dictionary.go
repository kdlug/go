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
