package dependency

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//changed fmt.Printf to --> fmt.Fprintf
//this is because fmt.Printf sends to stdout
//whereas fmt.Fprintf, will send the string to a Writer (anything that has the Write method)
