package app

import "github.com/boltdb/bolt"

func NewBoltStore(dbfile string) (*BoltDbStore, error) {
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltDbStore{
		db: db,
	}, nil
}

type BoltDbStore struct {
	db *bolt.DB
}

func (db *BoltDbStore) Init() {

}

func (db *BoltDbStore) Close() {

}

// AddMessage
// DeleteMessage
// GetMessages
// GetMessage
