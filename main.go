package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			exit(err)
		}
		w.WriteByte(corrupt(b))
		w.Flush()
	}
}

func corrupt(b byte) byte {
	pos := uint(random.Intn(8))
	return b ^ 1<<pos
}

func exit(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
