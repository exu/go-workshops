package api

import "fmt"

type ApiResourceCreator interface {
	Create()
}

type ApiResourceDeleter interface {
	Delete()
	IsDeleted() bool
}

type HttpResource struct {
	Name string
	Data map[string]interface{}
}

func (res *HttpResource) Create() {
	fmt.Println("creating:", res.Data)
}

func (res *HttpResource) Delete() {
	fmt.Println("deleting:", res.Data)
}

func BulkDeleter(resources []ApiResourceDeleter) {
	for _, resource := range resources {
		resource.Delete()
	}
}
