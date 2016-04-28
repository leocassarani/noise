package main

import (
	"io"
	"math/rand"
	"time"
)

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		out:    w,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type Writer struct {
	out    io.Writer
	random *rand.Rand
}

func (w *Writer) WriteByte(b byte) error {
	bytes := []byte{w.corrupt(b)}
	_, err := w.out.Write(bytes)
	return err
}

func (w *Writer) corrupt(b byte) byte {
	pos := uint(w.random.Intn(8))
	return b ^ 1<<pos
}
