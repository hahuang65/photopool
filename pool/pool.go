package pool

import (
	"time"

	"github.com/google/uuid"
)

type Pool struct {
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func New(name string) Pool {
	return Pool{
		Name:      name,
		ID:        uuid.New(),
		CreatedAt: time.Now().Truncate(0), // Get rid of monotonic clock
	}
}
