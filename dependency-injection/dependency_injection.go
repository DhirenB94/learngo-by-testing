package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Zaggy")
}

//here we have changed the greet function to have a Writer interface as the arguement
//now, as os.Stdout and bytes.Buffer bpth implement the io.Writer interface, both can be used in the greet funtion
