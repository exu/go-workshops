package handler

import (
	"net/http"

	"github.com/exu/go-workshops/090-basics-project-structure/lib/booer"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(booer.JSON))
}
