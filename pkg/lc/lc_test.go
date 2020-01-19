package lc

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestCountLines_SingleLine(t *testing.T) {
	sample := bytes.NewReader([]byte("Simple line\n"))

	got, _ := CountLines(sample)
	want := 2
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCountLines_SingleLineNoEndingBreak(t *testing.T) {
	sample := bytes.NewReader([]byte("Line1"))

	got, _ := CountLines(sample)
	want := 1
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCountLines_MultiLine(t *testing.T) {
	sample := bytes.NewReader([]byte("Line1\nLine2\nLine3\n"))

	got, _ := CountLines(sample)
	want := 4
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCountLines_MultiLineNoEndingBreak(t *testing.T) {
	sample := bytes.NewReader([]byte("Line1\nLine2"))

	got, _ := CountLines(sample)
	want := 2
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCountLines_EmptyFile(t *testing.T) {
	file, err := os.Open("testdata/empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	got, _ := CountLines(file)
	want := 1
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCountLines_OverBufferSize(t *testing.T) {
	file, err := os.Open("testdata/ipsum.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	got, _ := CountLines(file)
	want := 9
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}