package api

type ApiResourceDeleter interface {
	Delete()
	IsDeleted() bool
}

func BulkDeleter(resources []ApiResourceDeleter) {
	for _, resource := range resources {
		resource.Delete()
		println(resource.IsDeleted())
	}
}
