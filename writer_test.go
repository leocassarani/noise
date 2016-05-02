package main

import (
	"bytes"
	"testing"
)

func TestWriter_some(t *testing.T) {
	var buf bytes.Buffer

	w := NewWriter(&buf, 0.25)
	w.rng = &testRNG{
		seq: []float32{0.2, 0.5, 1, 0},
	}

	if err := w.WriteByte(0xCA); err != nil {
		t.Fatal(err)
	}

	b, err := buf.ReadByte()
	if err != nil {
		t.Fatal(err)
	}

	// We expect bits 0, 3, 4 and 7 to have been flipped.
	var expected byte = 0xCA ^ (1<<0 | 1<<3 | 1<<4 | 1<<7)
	if b != expected {
		t.Errorf("expected %#X, got %#X", expected, b)
	}
}

func TestWriter_all(t *testing.T) {
	var buf bytes.Buffer

	// A writer with an error rate of 1 will flip all bits.
	w := NewWriter(&buf, 1)

	in := []byte{0x00, 0xF0, 0xFF}
	if _, err := w.Write(in); err != nil {
		t.Fatal(err)
	}

	out := make([]byte, 3)
	_, err := buf.Read(out)
	if err != nil {
		t.Fatal(err)
	}

	// We expect all bits to have been flipped in every byte.
	expected := []byte{0xFF, 0x0F, 0x00}
	for i, b := range out {
		if b != expected[i] {
			t.Errorf("expected %#X, got %#X at index %d", expected[i], b, i)
		}
	}
}

func TestWriter_none(t *testing.T) {
	var buf bytes.Buffer

	// A writer with an error rate of 0 will never flip any bits.
	w := NewWriter(&buf, 0)

	if err := w.WriteByte(0xFF); err != nil {
		t.Fatal(err)
	}

	b, err := buf.ReadByte()
	if err != nil {
		t.Fatal(err)
	}

	// We expect b to have been unchanged.
	if b != 0xFF {
		t.Errorf("expected %#X, got %#X", 0xFF, b)
	}
}

type testRNG struct {
	seq   []float32
	index int
}

func (rng *testRNG) Float32() float32 {
	if len(rng.seq) == 0 {
		return 0
	}

	if rng.index >= len(rng.seq) {
		rng.index = 0
	}

	num := rng.seq[rng.index]
	rng.index++

	return num
}
