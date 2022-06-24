package bolt

import (
	"encoding/json"

	"git.sr.ht/~hwrd/photopool/pool"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

type Repository struct {
	db *bolt.DB
}

const bucket = "pools"

func New(db *bolt.DB) (*Repository, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	return &Repository{db: db}, err
}

func (r Repository) Create(p *pool.Pool) (pool.Pool, error) {
	poolB, err := json.Marshal(p)
	if err != nil {
		return *p, err
	}

	key, err := p.ID.MarshalBinary()
	if err != nil {
		return *p, err
	}

	return *p, r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		return b.Put(key, poolB)
	})
}

func (r Repository) Get(id uuid.UUID) (pool.Pool, error) {
	var p pool.Pool

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		key, err := id.MarshalBinary()
		if err != nil {
			return err
		}

		poolB := b.Get(key)
		return json.Unmarshal(poolB, &p)
	})

	return p, err
}
