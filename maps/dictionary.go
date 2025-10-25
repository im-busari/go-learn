package maps

// TODO: Read more https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps
type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// Map Lookup can return 2 values - value and success/ok.
	// The second value is a boolean which indicates if the key was found successfully.
	// It helps us differentiate between a key that doesn't exist in the map and a key that just doesn't
	// have a value/definition.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
	// So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the
	// underlying data structure that contains the data.
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
