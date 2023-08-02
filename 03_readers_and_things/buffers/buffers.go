package buffers

import "bytes"

type BufferWriter struct {
	buffer bytes.Buffer
}

func (bw *BufferWriter) Write(p []byte) (n int, err error) {
	return bw.buffer.Write(p)
}

func (bw *BufferWriter) String() string {
	return bw.buffer.String()
}
