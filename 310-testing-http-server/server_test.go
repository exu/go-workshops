package main

import (
	"./handler"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {
	// Used in E2E tests, configurable
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
			"Invalid greeting!\n I've got \t\t%s\n but want: \t%s",
			string(greeting),
			"Helloł Łerld",
		)
	}
}
