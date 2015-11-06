package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	buf := bytes.NewBuffer(nil)
	dir := "files/"
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		f, _ := os.Open(dir + file.Name())
		io.Copy(buf, f)
		f.Close()
	}
	s := string(buf.Bytes())

	println(s)
}
