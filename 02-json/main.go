package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string   `json:"mess"`
	To      []string `json:"recipients"`
	Code    int      `json:"secret_code"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		"Helloł Łerld",
		[]string{
			"Mom",
			"Dad",
			"Grandpa",
		},
		120,
	}

	data, _ := json.Marshal(response)

	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
