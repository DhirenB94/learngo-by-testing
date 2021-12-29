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

