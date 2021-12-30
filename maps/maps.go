package maps

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("the word you are looking for can not be found")
	ErrWordExists = DictionaryErr("the word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist so can not update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}


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


func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)

	//Go has a built-in function delete that works on maps.
	//It takes two arguments. The first is the map and the second is the key to be removed.
	//The delete function returns nothing, and we based our Delete method on the same notion.
	//Since deleting a value that's not there has no effect, unlike our Update and Add methods, we don't need to complicate the API with errors.
}