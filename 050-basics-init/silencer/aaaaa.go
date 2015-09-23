package silencer

import (
	"fmt"
	"io"
)

func init() {
	fmt.Println("Initialize silencer!!!")
}

func Whisp(writer io.Writer) {
	writer.Write([]byte("shhhhhhhhhh!"))
}
