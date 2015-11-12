// Demo of the expvar package. You register metrics by creating NewT, then
// update it.

// You can access the exposed metrics via HTTP at /debug/vars, you'll get a JSON
// object with your exposed variables and some pre defined system ones.

// You can use monitoring system such as Nagios and OpenNMS to monitor the
// system and plot the change of data over time.

// po uruchomieniu wywołaj "curl http://localhost:8080?user=placel"
// "curl http://localhost:8080/debug/vars | python -m json.tool".

// dodajmy interfejs do tych danych :)
// zainstaluj expvarmon
// $ go get github.com/divan/expvarmon
// $ expvarmon -ports="8080" -i 100ms -vars "rps,last_user"

package main

import (
	"expvar"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Two metrics, these are exposed by "magic" :)
// Number of calls to our server.
var rps = expvar.NewInt("rps")
var numCalls int64

// Last user.
var lastUser = expvar.NewString("last_user")

func HelloServer(w http.ResponseWriter, req *http.Request) {
	user := req.FormValue("user")

	numCalls++
	lastUser.Set(user)

	msg := fmt.Sprintf("G'day %s\n", user)
	io.WriteString(w, msg)
}

func stats() {
	for range time.Tick(time.Second) {
		rps.Set(numCalls)
		numCalls = 0
	}
}

func main() {
	go stats()
	http.HandleFunc("/", HelloServer)
	panic(http.ListenAndServe(":8080", nil))
}
