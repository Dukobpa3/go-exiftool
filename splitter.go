package exiftool

import (
	"bufio"
	"bytes"
	"io"
)

func splitReadyToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if bytes.Contains(data, endPattern) {
		endPos := bytes.Index(data, endPattern)
		if endPos > 0 {
			return endPos + len(endPattern), data[:endPos], bufio.ErrFinalToken
		}
	}

	if atEOF {
		return 0, data, io.EOF
	}
	return 0, nil, nil
}

func JsonSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if bytes.Contains(data, endPattern) {
		endPos := bytes.Index(data, endPattern)
		if endPos > 0 {
			return endPos + len(endPattern), data[:endPos], bufio.ErrFinalToken
		}
	}

	if i := bytes.IndexAny(data, "}]"); i >= 0 {
		return i + 1, data[:i+1], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, io.EOF

	}

	return 0, nil, nil
}

func RegularSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	fileDelimiter := []byte("======== ")

	if bytes.Contains(data, endPattern) {
		endPos := bytes.Index(data, endPattern)
		if endPos > 0 {
			return endPos + len(endPattern), data[:endPos], bufio.ErrFinalToken
		}
	}

	if i := bytes.Index(data, fileDelimiter); i >= 0 {
		if i > 0 {
			return i + len(fileDelimiter), data[:i], nil
		}
		return i + len(fileDelimiter), data[i:], nil
	}

	if atEOF && len(data) > 0 {
		return len(data), data, io.EOF

	}

	return 0, nil, nil
}
