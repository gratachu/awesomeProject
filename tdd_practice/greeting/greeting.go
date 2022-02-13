package greeting

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, word string) {
	fmt.Fprintf(writer, "Hello %s", word)
}
