package handler

import (
	"net/http"
)

func Simple(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(400)
	w.Write([]byte("Helloł Łerld"))
}
