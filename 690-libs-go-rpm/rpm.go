package main

import (
	"fmt"
	"github.com/cavaliercoder/go-rpm"
)

func main() {
	p, err := rpm.OpenPackageFile("p.rpm")
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Architecture())
	fmt.Println(p.ArchiveSize())
	fmt.Println(p.BuildHost())
	fmt.Println(p.BuildTime())
	fmt.Println(p.ChangeLog())
}
