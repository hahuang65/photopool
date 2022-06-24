package pool_test

import (
	"testing"

	"git.sr.ht/~hwrd/photopool/pool"
	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("AssignsName", func(t *testing.T) {
		want := "Mom's Retirement Party"
		p := pool.New(want)
		got := p.Name

		if want != got {
			t.Errorf("want: %q, got %q", want, got)
		}
	})

	t.Run("SetsID", func(t *testing.T) {
		p := pool.New("Richard's Graduation")

		if p.ID == uuid.Nil {
			t.Errorf("ID should be set, but is nil")
		}
	})

	t.Run("SetsCreatedAt", func(t *testing.T) {
		p := pool.New("Joey's Bar Mitzvah")

		if p.CreatedAt.IsZero() {
			t.Errorf("CreatedAt should be set, but is not")
		}
	})
}
