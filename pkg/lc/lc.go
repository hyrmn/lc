package lc

import (
	"bytes"
	"io"
)

// CountLines returns the number of lines in the provided byte stream
// It attempts to do this by counting occurances of `\n`
// If a file does not end with a `\n` then the last line of text will not be counted
// This is consistent with the behavior of running `wc -l <textfile.txt>`
func CountLines(r io.Reader) (int, error) {

	var count int
	var read int
	var err error
	const target byte = '\n'

	buffer := make([]byte, 32*1024)

	for {
		read, err = r.Read(buffer)
		if err != nil {
			break
		}

		var position int
		for {
			idxOf := bytes.IndexByte(buffer[position:read], target)
			if idxOf == -1 {
				break
			}

			count++
			position += idxOf + 1
		}
	}

	if err == io.EOF {
		return count, nil
	}

	return count, err
}
