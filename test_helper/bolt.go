package test_helper

import (
	"io/ioutil"
	"os"
	"testing"

	boltRepo "git.sr.ht/~hwrd/photopool/pool/repository/bolt"
	"github.com/boltdb/bolt"
)

func SetupBolt(t *testing.T) (*boltRepo.Repository, func()) {
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

	boltRepo, err := boltRepo.New(bolt)
	if err != nil {
		t.Fatalf("Could not create new Bolt repository: %v", err)
	}

	return boltRepo, teardown
}
