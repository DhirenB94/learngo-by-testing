package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("the word you are looking for can not be found")
	ErrWordExists = errors.New("the word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string)  error {
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
}

// I think....
//we are switching on the error
//d.search(word) retruns string and an error
// _ = string (we dont care about it) and err = error
//when you search an unknown word with Search(), you will return ErrNotFound as the error
//this mean the word does not already exist and you can add it --> adding to map --> d[word] = definition
//if the word does exist, Search() will return nil so we switch on this and return ErrWordExists as we dont want to add and overwrite
//Having a switch like this provides an extra safety net, in case Search() returns an error other than ErrNotFound