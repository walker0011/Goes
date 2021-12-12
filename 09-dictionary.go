package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for ")
var ErrWordExists = errors.New("cannot add word because it already exists")
var ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// 对于add来说，我们需要解决添加一寸的word时不改变原先word的definition
func (d Dictionary) Added(word, definition string) error {
	// 1. 添加之前，先查找看能够找到
	_, err := d.Search(word)

	switch err {
	// 2.1 找不到，则添加，并返回nil
	case ErrNotFound:
		d[word] = definition
	// 2.2 找到，则返回已存在error标志
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// 对于update来说，如果word不存在，我们不应该add一个
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

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
