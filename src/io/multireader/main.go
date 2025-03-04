package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("---- HEADER ----\n")
	body := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("---- FOOTER ----\n")

	reader := io.MultiReader(header, body, footer)
	io.Copy(os.Stdout, reader)
}
