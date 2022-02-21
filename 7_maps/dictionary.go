// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps

package maps

type Dictionary map[string]string // A map value is a pointer to a runtime.hmap structure, so when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.: A map value is a pointer to a runtime.hmap structure.

// cautions!
// You should never initialize an empty map variable, because maps can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. : https://go.dev/blog/maps
// `var m map[string]string`: NG!
// `var dictionary = map[string]string{}`: OK!
// `var dictionary = make(map[string]string)`: OK!

const ( // constant errors: https://dave.cheney.net/2016/04/07/constant-errors
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string { // implement 'error' interface. : https://go.dev/blog/error-handling-and-go
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) { // Since deleting a value that's not there has no effect, unlike our Update and Add methods, we don't need to complicate the API with errors.
	delete(d, word) // 'delete' is a built-in function: https://pkg.go.dev/builtin#delete
}
