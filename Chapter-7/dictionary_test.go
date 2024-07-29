package chapter7

import "testing"

func TestSearch(t *testing.T) {
	dictonary := Dictonary{"test": "this is just a test"}
	t.Run("kown word", func(t *testing.T) {
		got, _ := dictonary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unkown word", func(t *testing.T) {
		_, got := dictonary.Search("unkown")

		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictonary := Dictonary{}

	t.Run("Add should work", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictonary.Add(word, definition)

		assertDefinition(t, dictonary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictonary := Dictonary{word: definition}

		err := dictonary.Add(word, definition)

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictonary, word, definition)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("exsiting word", func(t *testing.T) {
		word := "test"
		oldDefinition := "this is a test"
		newDefinition := "this is not a test"

		dictonary := Dictonary{word: oldDefinition}

		err := dictonary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictonary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictonary := Dictonary{}

		err := dictonary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("exsisting word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictonary := Dictonary{word: definition}

		dictonary.Delete(word)

		_, err := dictonary.Search(word)
		assertError(t, err, ErrorNotFound)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictonary Dictonary, word, definition string) {
	t.Helper()

	got, err := dictonary.Search(word)

	if err != nil {
		t.Fatal("Should find word: ", err)
	}

	assertStrings(t, got, definition)
}
