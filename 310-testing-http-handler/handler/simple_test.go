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

	// We're injecting recorder to our handler instead
	// of writer injected by MUX
	Simple(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid code! I want %d but get %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Helloł Łerld" {
		t.Errorf("Invalid greeting, gimme Hello world phoneticaly in polish")
	}
}
