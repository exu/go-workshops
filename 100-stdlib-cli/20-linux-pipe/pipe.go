package main

import "os"
import "log"
import "io/ioutil"
import "fmt"

func main() {

	stat, _ := os.Stdin.Stat()

	bytes, err := ioutil.ReadAll(os.Stdin)
	log.Println(err, string(bytes))
}
