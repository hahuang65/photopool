package service_test

import (
	"testing"

	"git.sr.ht/~hwrd/photopool/pool/service"
	helper "git.sr.ht/~hwrd/photopool/test_helper"
)

func TestCreateWithName(t *testing.T) {
	t.Parallel()

	b, teardown := helper.SetupBolt(t)
	defer teardown()

	cases := []struct {
		name    string
		service *service.Service
	}{
		{"BoltDB", service.New(b)},
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

	b, teardown := helper.SetupBolt(t)
	defer teardown()

	cases := []struct {
		name    string
		service *service.Service
	}{
		{"BoltDB", service.New(b)},
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
