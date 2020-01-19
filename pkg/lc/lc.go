package lc

import (
	"bytes"
	"io"
)

// CountLines returns the number of lines in the provided byte stream
// It attempts to do this by counting occurances of `\n` while also assuming
// that any input read counts as at least one line. This handles cases where a file may not end with a `\n`
func CountLines(r io.Reader) (int, error) {

	var count int
	var read int
	var err error
	var hasContent bool

	buffer := make([]byte, 32*1024)
	hasContent = false

	for {
		read, err = r.Read(buffer)
		if err != nil {
			break
		}

		hasContent = true

		var position int
		for {
			idxOf := bytes.IndexByte(buffer[position:read], '\n')
			if idxOf == -1 {
				break
			}

			count++
			position += idxOf + 1
		}
	}

	if err == io.EOF {
		if hasContent || count == 0 {
			count++
		}
		return count, nil
	}

	return count, err
}
