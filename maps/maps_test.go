package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test" : "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want )
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "the word you are looking for can not be found"

		if err == nil {
			t.Fatalf("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string)  {
	t.Helper()

	if got!= want {
		t.Errorf("got %q, want %q", got, want)
	}
}

//what happens if the word we search for is not in the map - we want to communicate this to the user
//we return a 2nd argument which is an Error type
//Errors can be converted to a string with the .Error() method, which we do when passing it to the assertion
//We are also protecting assertStrings with if to ensure we don't call .Error() on nil.


//current test output
// ./maps_test.go:16:10: assignment mismatch: 2 variables but dictionary.Search returns 1 value