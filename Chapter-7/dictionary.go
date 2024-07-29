package chapter7

import "errors"

type Dictonary map[string]string

var (
	ErrorNotFound         = errors.New("could not find the word you were looking for")
	ErrorWordExists       = errors.New("word has already been added to the dictonary")
	ErrorWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictonary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictonary) Add(word, definition string) error {
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

func (d Dictonary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err

	}

	return nil
}

func (d Dictonary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return err
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
