package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)	// menulis di terminal
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("intan sukma\n")

	writer.Flush()

}