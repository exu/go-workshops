package package2

import (
	"fmt"
	"io"
)

func init() {
	fmt.Println("Initialize package2!!!")
}

func Write(writer io.Writer) {
	writer.Write([]byte("package2: shhhhhhhhhh!\n"))
}
