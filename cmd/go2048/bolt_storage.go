package main

import (
	"fmt"

	"github.com/boltdb/bolt"

	"github.com/gitchander/go2048/core"
)

type BoltStorage struct {
	db         *bolt.DB
	bucketName []byte
}

var _ core.Storage = &BoltStorage{}

func NewBoltStorage(filename string) (*BoltStorage, error) {

	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}

	bucketName := []byte("go2048")

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			_, err := tx.CreateBucket(bucketName)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &BoltStorage{
		db:         db,
		bucketName: bucketName,
	}, nil
}

func (st *BoltStorage) Close() error {
	if st.db == nil {
		return nil
	}
	err := st.db.Close()
	st.db = nil
	return err
}

func (st *BoltStorage) Put(key string, value []byte) error {
	return st.db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket(st.bucketName)
			if b == nil {
				return fmt.Errorf("no bucket %s in db", st.bucketName)
			}
			return b.Put([]byte(key), value)
		},
	)
}

func (st *BoltStorage) Get(key string) (value []byte, err error) {
	err = st.db.View(
		func(tx *bolt.Tx) error {
			b := tx.Bucket(st.bucketName)
			if b == nil {
				return fmt.Errorf("no bucket %s in db", st.bucketName)
			}
			value = b.Get([]byte(key))
			return nil
		},
	)
	return value, err
}

func (st *BoltStorage) Remove(key string) error {
	return st.db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket(st.bucketName)
			if b == nil {
				return fmt.Errorf("no bucket %s in db", st.bucketName)
			}
			return b.Delete([]byte(key))
		},
	)
}
