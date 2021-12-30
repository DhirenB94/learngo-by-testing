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

	t.Run("add already existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
	 //Current errors
	//maps_test.go:44: got error %!q(<nil>), want "the word already exists"
	//maps_test.go:45: got "new test", want "this is just a test"
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


