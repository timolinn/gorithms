package datastructures

import "unsafe"

type StringBuilder struct {
	buf []byte
}

// WriteString writes to the data buffer
func (sb *StringBuilder) WriteString(s string) {
	sb.buf = append(sb.buf, s...)
}

func (sb *StringBuilder) String() string {
	// return string(sb.buf)
	return *(*string)(unsafe.Pointer(&sb.buf))
}
