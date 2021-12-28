package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}

//here we have creates a Dictionary type which acts as a thin wrapper around map
//then creates a Search method