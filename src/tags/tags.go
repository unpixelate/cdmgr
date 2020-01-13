package tags

import (
	"../db"
	"../initialiser"
	"github.com/boltdb/bolt"
)

func Init(taskBucket string) error {
	if taskBucket == initialiser.TagMasterData {
		return db.RaiseErr("Cannot modifiy master db", taskBucket)
	}
	bucket := []byte(taskBucket)
	var err error
	if err != nil {
		return err
	}
	return db.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
}
