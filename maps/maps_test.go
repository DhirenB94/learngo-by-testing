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
//new test to add words
//utilise our search method to make validation of our dictionary easier

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}


