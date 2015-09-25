package screamer

import (
	"fmt"
	"github.com/exu/go-workshops/041-basics-init/silencer"
	"io"
)

func init() {
	fmt.Println("Initialize screamer!!!")
}

func ScreamAAAA(writer io.Writer) {
	writer.Write([]byte("aaaaaaaaa"))
	silencer.Whisp(writer)
}
