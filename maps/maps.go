package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}

//as we are returning nil for the error value our test instructs it to t.fatal with "expected to get an error"