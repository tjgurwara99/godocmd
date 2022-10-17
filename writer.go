package main

import (
	"bytes"
	"io"
)

type writer struct {
	writer io.Writer
	prefix string
}

// NewWriter Creates a prefix writer such that all
// new lines are prefixed with the given prefix.
func NewWriter(old io.Writer, prefix string) io.Writer {
	return &writer{
		writer: old,
		prefix: prefix,
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	n = len(p)
	var buf bytes.Buffer
	for _, line := range bytes.Split(p, []byte("\n")) {
		if bytes.TrimSpace(bytes.ReplaceAll(line, []byte{'\n'}, []byte{' '})) != nil {
			buf.WriteString(w.prefix)
		}
		buf.Write(line)
		buf.WriteString("\n")
	}
	buf.Truncate(buf.Len() - 1)
	_, err = w.writer.Write(buf.Bytes())
	return n, err
}
