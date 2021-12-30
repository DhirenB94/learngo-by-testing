package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test" : "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want )
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word:= "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil) //modified to check for a nil error
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("trying to add an already existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

//update the definition of an existing word

func TestUpdate(t *testing.T) {
	t.Run("existing word so update definition", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"

		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)	
	})
	t.Run("new word so can not update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
	// Current errors :
	// ./maps_test.go:60:27: dictionary.Update(word, newDefinition) used as value
	// ./maps_test.go:71:27: dictionary.Update(word, definition) used as value
	// ./maps_test.go:73:23: undefined: ErrWordDoesNotExist
	

}

func assertStrings(t testing.TB, got, want string)  {
	t.Helper()

	if got!= want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got!= want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string)  {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}
}


