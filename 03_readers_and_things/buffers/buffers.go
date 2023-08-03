package buffers

import (
	"bytes"
	"io"
)

type BufferWriter struct {
	buffer bytes.Buffer
}

func NewBufferWriter() *BufferWriter {
	return &BufferWriter{
		buffer: bytes.Buffer{},
	}
}

func (bw *BufferWriter) Write(p []byte) (n int, err error) {
	return bw.buffer.Write(p)
}

func (bw *BufferWriter) String() string {
	return bw.buffer.String()
}

type BufferReader struct {
	r io.Reader
}

func NewBufferReader(r io.Reader) *BufferReader {
	return &BufferReader{
		r: r,
	}
}

func (br *BufferReader) Read(p []byte) (n int, err error) {
	return br.r.Read(p)
}
