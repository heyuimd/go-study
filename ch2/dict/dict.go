package dict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("not found")
	errWordExists = errors.New("already exists")
)

func (d *Dictionary) Search(word string) (string, error) {
	value, ok := (*d)[word]

	if ok {
		return value, nil
	}

	return "", errNotFound
}

func (d *Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		(*d)[word] = def
		return nil
	}

	return errWordExists
}
