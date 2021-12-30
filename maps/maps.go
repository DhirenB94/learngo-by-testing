package maps

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("the word you are looking for can not be found")
	ErrWordExists = DictionaryErr("the word already exists")
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
