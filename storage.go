package go2048

import "errors"

type Storage interface {
	Put(key string, val []byte) error
	Get(key string) (val []byte, err error)
	Remove(key string) error
}

type FakeStorage map[string][]byte

var _ Storage = &FakeStorage{}

func NewFakeStorage() *FakeStorage {
	fs := FakeStorage(make(map[string][]byte))
	return &fs
}

func (fs *FakeStorage) Put(key string, val []byte) error {
	(*fs)[key] = val
	return nil
}

func (fs *FakeStorage) Get(key string) (val []byte, err error) {
	val, ok := (*fs)[key]
	if !ok {
		return nil, errors.New("has not key in storage")
	}
	return val, nil
}

func (fs *FakeStorage) Remove(key string) error {
	_, ok := (*fs)[key]
	if !ok {
		return errors.New("has not key in storage")
	}
	delete((*fs), key)
	return nil
}
