package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB
var errInit error

type Dirs struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	DB, errInit = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if errInit != nil {
		return errInit
	}
	return nil
}

func CreateBucket(taskBucket string) error {
	Bucket := []byte(taskBucket)
	var err error
	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(Bucket)
		return err
	})
	return err
}

func CreateDirectory(dir string, taskBucket string) (int, error) {
	Bucket := []byte(taskBucket)
	var id int
	err := DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(dir))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllDir(taskBucket string) ([]Dirs, error) {
	Bucket := []byte(taskBucket)
	var dirs []Dirs
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			dirs = append(dirs, Dirs{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

func DeleteDir(key int, taskBucket string) error {
	Bucket := []byte(taskBucket)
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		return b.Delete(itob(key))
	})
}

func DeleteBucket(taskBucket string) error {
	Bucket := []byte(taskBucket)
	return DB.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(Bucket)
		return err
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
