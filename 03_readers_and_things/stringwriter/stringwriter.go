package stringwriter

type StringConcatWriter struct {
	data string
}

func (w *StringConcatWriter) Write(p []byte) (n int, err error) {
	w.data += string(p)
	return len(p), nil
}

func (w *StringConcatWriter) String() string {
	return w.data
}
