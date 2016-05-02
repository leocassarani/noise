package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	errorRate := flag.Float64("rate", 0.25, "error rate, must be in the [0,1) range")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	w := NewWriter(os.Stdout, *errorRate)

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
