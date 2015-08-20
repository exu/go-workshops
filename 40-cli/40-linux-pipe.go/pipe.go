package main

import "os"
import "log"
import "io/ioutil"
import "fmt"

func main() {

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		fmt.Println("data is being piped to stdin")
		bytes, err := ioutil.ReadAll(os.Stdin)
		log.Println(err, string(bytes))
	} else {
		fmt.Println("stdin is from a terminal")
	}

}
