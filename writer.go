package main

import (
	"io"
	"math/rand"
	"time"
)

func NewWriter(out io.Writer, errorRate float32) *Writer {
	return &Writer{
		out:       out,
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())),
		errorRate: errorRate,
	}
}

type Writer struct {
	errorRate float32
	out       io.Writer
	rng       rng
}

func (w *Writer) Write(p []byte) (n int, err error) {
	for _, b := range p {
		err = w.WriteByte(b)
		if err != nil {
			break
		}
		n++
	}
	return n, err
}

func (w *Writer) WriteByte(b byte) error {
	for i := uint(0); i < 8; i++ {
		if w.rng.Float32() < w.errorRate {
			b ^= 1 << i
		}
	}

	_, err := w.out.Write([]byte{b})
	return err
}

// RNG = Random Number Generator
type rng interface {
	Float32() float32
}
