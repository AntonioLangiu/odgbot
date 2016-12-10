package bot

import (
	"log"
	"github.com/boltdb/bolt"
	"github.com/AntonioLangiu/odgbot/common"
)

func LoadDb (configuration *common.Configuration) *bolt.DB{
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
        log.Fatal(err)
    }
	return db
}

func CreateUser (db *bolt.DB, username string) error {
	return db.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket([]byte(username))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
}
