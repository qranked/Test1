package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

func initBoltdb(path string, mode os.FileMode) (*bolt.DB, error) {
	db, err := bolt.Open("main.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Error occured on open Boltdb db: %s", err)
	}
	return db, nil
}

func createBoltdbBucket(names ...string) error {
	return db.Update(func(tx *bolt.Tx) error {
		var err error
		for _, name := range names {
			_, err = tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				break
			}
		}
		if err != nil {
			return fmt.Errorf("Error occured on create Boltdb bucket: %s", err)
		}
		return nil
	})
}

func setBoltValue(bucket string, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(value))
		return err
	})
}

func getBoltRawValue(bucket, key string) []byte {
	var value []byte
	db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte(bucket)).Get([]byte(key))
		return nil
	})
	return value
}

func getBoltStrValue(bucket, key string) string {
	var value string
	db.View(func(tx *bolt.Tx) error {
		value = string(tx.Bucket([]byte(bucket)).Get([]byte(key)))
		return nil
	})
	return value
}
