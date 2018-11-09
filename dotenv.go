package dotenv

import "bufio"

// Reader to capture properties of Dot env Reader
type Reader struct {
	// tracks the line number
	lineNum string
	// buffer of current line
	buffer []byte
	// reader
	r *bufio.Reader
}
