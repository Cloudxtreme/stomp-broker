package main

import (
	"bytes"
)

// dropNULL drops a NULL \0 from the data.
func dropNULL(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\x00' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanSTOMPBody(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\x00'); i >= 0 {
		// We have a non-empty body.
		return i + 1, dropNULL(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropNULL(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
