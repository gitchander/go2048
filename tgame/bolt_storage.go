package main

import (
	"fmt"

	"github.com/boltdb/bolt"

	game "github.com/gitchander/go2048"
)

type BoltStorage struct {
	db         *bolt.DB
	bucketName []byte
}

var _ game.Storage = &BoltStorage{}

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

func (st *BoltStorage) Put(key string, val []byte) error {
	return st.db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket(st.bucketName)
			if b == nil {
				return fmt.Errorf("no bucket %s in db", st.bucketName)
			}
			return b.Put([]byte(key), val)
		},
	)
}

func (st *BoltStorage) Get(key string) (val []byte, err error) {
	err = st.db.View(
		func(tx *bolt.Tx) error {
			b := tx.Bucket(st.bucketName)
			if b == nil {
				return fmt.Errorf("no bucket %s in db", st.bucketName)
			}
			val = b.Get([]byte(key))
			return nil
		},
	)
	return val, err
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
