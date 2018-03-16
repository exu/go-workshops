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

	// we want to test if our integration works well,
	// it'll be valid when Some deleter will call Delete on all objects.
	// we're passing our mock to BulkDeleter

	for i, res := range ress {
		if !res.IsDeleted() {
			t.Errorf("Resource `%d` should be deleted", i)
		}
	}
}
