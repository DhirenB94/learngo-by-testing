package maps


func Search(dictionary map[string]string, word string) string{
	return dictionary[word]
}
//getting a value out of map is same as getting one from an array
//map[key]