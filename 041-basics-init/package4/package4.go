package package4

import (
	"fmt"
	"io"
)

func init() {
	fmt.Println("Initialize package4!!!")
}

func Write(writer io.Writer) {
	writer.Write([]byte("package4: wrrrrrrrrrrrrr!\n"))
}
