package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("the word you are looking for can not be found")
	}

	return definition, nil
}

//In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values.
//The second value is a boolean which indicates if the key was found successfully.
//This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.