package bolt_test

import (
	"testing"

	"git.sr.ht/~hwrd/photopool/pool"
	helper "git.sr.ht/~hwrd/photopool/test_helper"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	r, teardown := helper.SetupBolt(t)
	defer teardown()

	p := pool.New("Foo")

	t.Run("ReturnsNilForSuccess", func(t *testing.T) {
		_, err := r.Create(&p)
		assert.Nil(t, err)
	})

	t.Run("PersistsAPool", func(t *testing.T) {
		want, _ := r.Create(&p)
		got, _ := r.Get(p.ID)
		assert.Equal(t, want, got)
	})
}

func TestGet(t *testing.T) {
	t.Parallel()

	r, teardown := helper.SetupBolt(t)
	defer teardown()

	p := pool.New("Foo")

	want, err := r.Create(&p)
	if err != nil {
		t.Fatalf("Could not create pool: %v", err)
	}

	t.Run("GetsAPool", func(t *testing.T) {
		got, err := r.Get(p.ID)
		if err != nil {
			t.Fatalf("Could not get pool with ID %s: %v", p.ID, err)
		}

		assert.Equal(t, want, got)
	})
}
