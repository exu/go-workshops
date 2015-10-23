package package1

import (
	"fmt"
	"io"
)

func init() {
	fmt.Println("Initialize pacakge1!!!")
}

func Write(writer io.Writer) {
	writer.Write([]byte("package1: aaaaaaaaa\n"))
}
