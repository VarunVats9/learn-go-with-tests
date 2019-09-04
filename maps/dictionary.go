package maps

const (
	// ErrorNotFound specifies the error in case the key is not found.
	ErrorNotFound = DictionaryErr("could not find the word you are looking for")
	// ErrorWordExists specifies the error in case the key already exists.
	ErrorWordExists = DictionaryErr("cannot add word as it already exists")
	// ErrorWordDoesNotExists specifies the error in case the key does not exist, while updating
	ErrorWordDoesNotExists = DictionaryErr("cannot update word as it does not exist")
)

// DictionaryErr is a type which implements error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary is a thin wrapper over map
type Dictionary map[string]string

// Search takes dictionary, key and returns the value
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

// Add takes key, value pair and saves it
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

// Update takes key, and new value and replace the old value
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

// Delete takes key, and removes it
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
