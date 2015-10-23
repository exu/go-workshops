package package3

import (
	"fmt"
	"github.com/exu/go-workshops/041-basics-init/package1"
	"github.com/exu/go-workshops/041-basics-init/package2"
	"github.com/exu/go-workshops/041-basics-init/package4"
	"io"
)

func init() {
	fmt.Println("Initialize package3!!!")
}

func Write(writer io.Writer) {
	package1.Write(writer)
	package2.Write(writer)
	package4.Write(writer)
}
