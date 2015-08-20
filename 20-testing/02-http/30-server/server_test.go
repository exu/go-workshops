package main

import (
	"./handler" // normalnie używamy ścieżek "Github" style

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {
	// Używany w testach E2E, konfigurowalny
	testServer := httptest.NewServer(http.HandlerFunc(handler.Simple))
	defer testServer.Close()

	response, err := http.Get(testServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(greeting) != "Helloł Łerld" {
		t.Errorf(
			"Nikt się nie przywitał :( lub nie to pozdrowienie!\n Jest \t\t%s\n miało być: \t%s",
			string(greeting),
			"Helloł Łerld",
		)
	}
}
