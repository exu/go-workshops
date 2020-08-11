package handler

import (
	"net/http"

	"myapp/lib/booer"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(booer.JSON))
}
