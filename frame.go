package main

import (
	"bufio"
	"net"
	"strings"
)

type Frame struct {
	Command string
	Headers Headers
	Body    string
}

func ReadFrame(conn net.Conn) (*Frame, error) {
	scanner := bufio.NewScanner(conn)

	// Get the STOMP command.
	if res := scanner.Scan(); !res {
		return nil, scanner.Err()
	}
	command := scanner.Text()

	// Get the STOMP Headers.
	headers := make(Headers, 0)
	for scanner.Scan() {
		rawHeader := scanner.Text()
		if len(rawHeader) == 0 {
			break
		}
		headerSlice := strings.Split(rawHeader, ":")
		header := Header{
			Key:   headerSlice[0],
			Value: headerSlice[1],
		}
		headers = append(headers, header)
	}

	// Get the STOMP Body.
	scanner.Split(ScanSTOMPBody)
	if res := scanner.Scan(); !res {
		return nil, scanner.Err()
	}
	body := scanner.Text()

	frame := &Frame{
		Command: command,
		Headers: headers,
		Body:    body,
	}
	return frame, nil
}
