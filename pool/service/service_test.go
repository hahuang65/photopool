package service_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"

	boltRepo "git.sr.ht/~hwrd/photopool/pool/repository/bolt"
	"git.sr.ht/~hwrd/photopool/pool/service"
)

func setupBolt(t *testing.T) (*bolt.DB, func()) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}

	bolt, err := bolt.Open(f.Name(), 0600, nil)
	if err != nil {
		t.Fatalf("Could not open BoltDB: %v", err)
	}

	teardown := func() {
		bolt.Close()
		f.Close()
		os.Remove(f.Name())
	}

	return bolt, teardown
}

func TestCreateWithName(t *testing.T) {
	t.Parallel()

	b, teardown := setupBolt(t)
	defer teardown()

	boltRepo, err := boltRepo.New(b)
	if err != nil {
		t.Fatalf("Could not create new Bolt repository: %v", err)
	}

	cases := []struct {
		name    string
		service *service.Service
	}{
		{"BoltDB", service.New(boltRepo)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("PersistsAPool", func(t *testing.T) {
				want, err := c.service.CreateWithName("Foo")
				if err != nil {
					t.Fatalf("Could not create pool: %v", err)
				}

				got, err := c.service.GetByID(want.ID)
				if err != nil {
					t.Fatalf("Could not get pool: %v", err)
				}

				if want != got {
					t.Errorf("want: %+v, got: %+v", want, got)
				}
			})
		})
	}
}

func TestGetByID(t *testing.T) {
	t.Parallel()

	b, teardown := setupBolt(t)
	defer teardown()

	boltRepo, err := boltRepo.New(b)
	if err != nil {
		t.Fatalf("Could not create new Bolt repository: %v", err)
	}

	cases := []struct {
		name    string
		service *service.Service
	}{
		{"BoltDB", service.New(boltRepo)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want, err := c.service.CreateWithName("Foo")
			if err != nil {
				t.Fatalf("Could not create pool: %v", err)
			}

			t.Run("GetsAPool", func(t *testing.T) {
				got, err := c.service.GetByID(want.ID)
				if err != nil {
					t.Fatalf("Could not get pool: %v", err)
				}

				if want != got {
					t.Errorf("want: %+v, got: %+v", want, got)
				}
			})
		})
	}
}
