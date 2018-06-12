package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		if strings.HasSuffix(file.Name(), ".go") {
			dirName := strings.Replace(file.Name(), ".go", "", 1)
			err := os.Mkdir(dirName, 0770)
			if err != nil {
				panic(err)
			}

			err = os.Rename(file.Name(), fmt.Sprintf("%s/%s", dirName, file.Name()))
			if err != nil {
				panic(err)
			}
		}
	}
}
