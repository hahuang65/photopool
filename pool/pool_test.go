package pool_test

import (
	"testing"

	"git.sr.ht/~hwrd/photopool/pool"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("AssignsName", func(t *testing.T) {
		want := "Mom's Retirement Party"
		p := pool.New(want)
		got := p.Name

		assert.Equal(t, want, got)
	})

	t.Run("SetsID", func(t *testing.T) {
		p := pool.New("Richard's Graduation")

		assert.NotNil(t, p.ID)
	})

	t.Run("SetsCreatedAt", func(t *testing.T) {
		p := pool.New("Joey's Bar Mitzvah")

		assert.NotNil(t, p.CreatedAt)
	})
}
