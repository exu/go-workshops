package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "http://l:8080/api", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()

	// Wstrzykujemy recorder (zamiast writera który jest normalnie
	// wstrzykiwany przez Mux) oraz stworzony request
	Simple(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Nie ten kod! Chciałem %d a mam %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Helloł Łerld" {
		t.Errorf("Nikt się nie przywitał :( lub nie to pozdrowienie!")
	}
}
