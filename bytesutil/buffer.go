package bytesutil

import "io"

// Buffer is a simple buffer for bytes.
// Only one of Read or Write methods should be used for the same buffer.
type Buffer struct {
	buf  []byte
	pos  int
	auto bool
}

// NewBuffer creates a new buffer with the given byte slice.
func NewBuffer(buf []byte, autoResize bool) *Buffer {
	return &Buffer{buf: buf, auto: autoResize}
}

// Buf returns the byte slice of buffer.
func (self *Buffer) Buf() []byte {
	return self.buf
}

// Pos returns the current position of buffer (i.e. total bytes read or written).
func (self *Buffer) Pos() int {
	return self.pos
}

// Read reads up to len(p) bytes from buffer into p.
// If auto is true and the buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read.
// If auto is false and the buffer does not have enough bytes to read, it will read available bytes and return the number of bytes read and error.
// If there is enough bytes to read, it will read len(p) bytes and return the number of bytes read.
func (self *Buffer) Read(p []byte) (int, error) {
	n := len(p)
	l := self.pos + n
	b := len(self.buf)
	if l <= b {
		i := copy(p, self.buf[self.pos:l])
		self.pos += i
		return n, nil
	}

	n = copy(p, self.buf[self.pos:])
	self.pos += n
	if self.auto {
		return n, nil
	}
	return n, io.ErrShortBuffer
}

// Write writes up to len(p) bytes from p to buffer.
// If auto is true and the buffer does not have enough space to write, it will resize the buffer and write p to buffer.
// If auto is false and the buffer does not have enough space to write, it will write available space and return the number of bytes written and error.
// If there is enough space to write, it will write len(p) bytes and return the number of bytes written.
func (self *Buffer) Write(p []byte) (int, error) {
	n := len(p)
	l := self.pos + n
	b := len(self.buf)
	if l <= b {
		i := copy(self.buf[self.pos:], p)
		self.pos += i
		return n, nil
	}

	if self.pos < b {
		i := copy(self.buf[self.pos:], p[:b-self.pos])
		if self.auto {
			self.buf = append(self.buf, p[i:]...)
			self.pos = len(self.buf)
			return n, nil
		}
		self.pos += i
		return i, io.ErrShortWrite
	}

	if self.auto {
		self.buf = append(self.buf, p...)
		self.pos = len(self.buf)
		return n, nil
	}

	return 0, io.ErrShortWrite
}
