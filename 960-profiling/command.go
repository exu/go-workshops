package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func init() {
	flag.Parse()
	cpuProfile(*cpuprofile)
}

func cpuProfile(fileName string) {
	if fileName != "" {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
}

func memProfile(fileName string) {
	if fileName != "" {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}

// Read something
func Read() string {
	content, _ := ioutil.ReadFile("file.txt")
	return string(content)
}

func main() {
	defer pprof.StopCPUProfile()
	defer memProfile(*memprofile)

	data := ""
	for i := 0; i < 100; i++ {
		data += Read()
	}
}
