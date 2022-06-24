package repository

import (
	"git.sr.ht/~hwrd/photopool/pool"
	"github.com/google/uuid"
)

type Repository interface {
	Create(*pool.Pool) (pool.Pool, error)
	Get(uuid.UUID) (pool.Pool, error)
}
