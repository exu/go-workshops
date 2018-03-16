package main

import (
	"encoding/json"
	"net/http"
)

// definiujemy struckturę z "annotacjami konfigurujacymi zwracane
// pola w JSONie
type Response struct {
	Message string   `json:"mess"`
	To      []string `json:"recipients"`
	Code    int      `json:"secret_code"`
}

// prosty HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {

	// nowa obiekt
	response := Response{
		"Helloł Łerld",
		[]string{
			"Mom",
			"Dad",
			"Grandpa",
		},
		120,
	}

	// marshalling
	data, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")

	// strumieniujemy dane
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}
