package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := NewWriter(os.Stdout, 0.25)

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			exit(err)
		}
		if err := w.WriteByte(b); err != nil {
			exit(err)
		}
	}
}

func exit(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
