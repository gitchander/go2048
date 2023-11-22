package core

import "errors"

var ErrNoValueByKey = errors.New("no value by key in storage")

type Storage interface {
	Put(key string, value []byte) error
	Get(key string) (value []byte, err error)
	Remove(key string) error
}

type MapStorage map[string][]byte

var _ Storage = &MapStorage{}

func NewMapStorage() *MapStorage {
	ms := MapStorage(make(map[string][]byte))
	return &ms
}

func (ms *MapStorage) Put(key string, val []byte) error {
	(*ms)[key] = val
	return nil
}

func (ms *MapStorage) Get(key string) (val []byte, err error) {
	val, ok := (*ms)[key]
	if !ok {
		return nil, ErrNoValueByKey
	}
	return val, nil
}

func (ms *MapStorage) Remove(key string) error {
	_, ok := (*ms)[key]
	if !ok {
		return ErrNoValueByKey
	}
	delete((*ms), key)
	return nil
}
