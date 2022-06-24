package bolt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"

	"git.sr.ht/~hwrd/photopool/pool"
	boltRepo "git.sr.ht/~hwrd/photopool/pool/repository/bolt"
)

func setup(t *testing.T) (*boltRepo.Repository, func()) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Could not create temp file: %v", err)
	}

	b, err := bolt.Open(f.Name(), 0600, nil)
	if err != nil {
		t.Errorf("Could not open file for BoltDB: %v", err)
	}

	teardown := func() {
		b.Close()
		f.Close()
		os.Remove(f.Name())
	}

	repo, err := boltRepo.New(b)
	if err != nil {
		t.Errorf("Could not create BoltDB repository: %v", err)
	}

	return repo, teardown
}

func TestCreate(t *testing.T) {
	t.Parallel()

	r, teardown := setup(t)
	defer teardown()

	p := pool.New("Foo")

	t.Run("ReturnsNilForSuccess", func(t *testing.T) {
		_, err := r.Create(&p)
		if err != nil {
			t.Error("Successful creation should return nil, but did not")
		}
	})

	t.Run("PersistsAPool", func(t *testing.T) {
		want, _ := r.Create(&p)
		got, _ := r.Get(p.ID)
		if got != want {
			t.Errorf("got: %+v, want: %+v", got, p)
		}
	})
}

func TestGet(t *testing.T) {
	t.Parallel()

	r, teardown := setup(t)
	defer teardown()

	p := pool.New("Foo")

	want, err := r.Create(&p)
	if err != nil {
		t.Errorf("Could not create pool: %v", err)
	}

	t.Run("GetsAPool", func(t *testing.T) {
		got, err := r.Get(p.ID)
		if err != nil {
			t.Errorf("Could not get pool with ID %s: %v", p.ID, err)
		}

		if want != got {
			t.Errorf("got: %+v, want: %+v", got, p)
		}
	})
}
