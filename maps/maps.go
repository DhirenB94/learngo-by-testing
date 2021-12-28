package maps

type Dictionary map[string]string


func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}

//it fails in the unknown word test because you are searching "unkwown"key in dictionary
//dictionary only has "test"key --> so it returns nil and therefore t.fatal