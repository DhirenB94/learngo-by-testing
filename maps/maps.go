package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("the word you are looking for can not be found")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
//adding to a map is similar to an array
//need to specify a key and set it equal to a value
// d[test] = this is just a test
