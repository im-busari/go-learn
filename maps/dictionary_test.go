package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("get definition of test", func(t *testing.T) {
		keyword := "test"
		dictionary := Dictionary{keyword: "this is just a test"}

		got, _ := dictionary.Search(keyword)
		want := "this is just a test"

		assertStrings(t, got, want, keyword)
	})

	t.Run("when keyword missing, get an empty string", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("missing")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "whale"
		definition := "A large animal in the sea"
		err := dictionary.Add(word, definition)

		if err != nil {
			t.Error("didn't not expected an error, but got one")
		}

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("throw an error when word exists", func(t *testing.T) {
		word := "whale"
		definition := "is just swimming"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "not swimming")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "otter"

		dictionary := Dictionary{word: "I am cute"}
		err := dictionary.Update(word, "I am hungry")

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, "I am hungry")
	})

	t.Run("return error if word is missing", func(t *testing.T) {
		dictionary := Dictionary{"whale": "I am cute"}
		err := dictionary.Update("otter", "I am hungry")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing item", func(t *testing.T) {
		word := "whale"
		dictionary := Dictionary{word: "I am swimming"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("delete non-existing item", func(t *testing.T) {
		word := "whale"
		dictionary := Dictionary{word: "I am swimming"}

		err := dictionary.Delete("otter")

		assertError(t, err, ErrWordDoesNotExist)

		_, err = dictionary.Search(word)

		assertError(t, err, nil)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want, keyword string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but want %q given, %q", got, want, keyword)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition, word)
}
