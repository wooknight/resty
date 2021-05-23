package data

import "errors"

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("No Such Key")

func Put(key, value string) (err error) {

	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	val, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return val, nil
}

func Delete(key string) (err error) {
	delete(store, key)
	return
}
