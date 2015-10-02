package api

import (
	"testing"
)

type MyResource struct {
	deleted bool
}

func (res *MyResource) Delete() {
	res.deleted = true
}

func (res MyResource) IsDeleted() bool {
	return res.deleted
}

func TestBulkDelete_ShouldCallDeleteOnAllResources(t *testing.T) {

	ress := []ApiResourceDeleter{
		&MyResource{},
		&MyResource{},
		&MyResource{},
	}

	BulkDeleter(ress)

	for i, res := range ress {
		if !res.IsDeleted() {
			t.Errorf("Resource `%d` should be deleted", i)
		}
	}
}
