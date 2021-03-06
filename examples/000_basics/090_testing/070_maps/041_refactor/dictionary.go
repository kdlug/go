// define custom error type

package main

func main() {

}

type Dictionary map[string]string

// We made the errors constant; this required us to create our own DictionaryErr type which implements the error interface.
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		//return "", errors.New("could not find the word you were looking for")
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type.
	// A gotcha that reference types introduce is that maps can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

	return nil
}
