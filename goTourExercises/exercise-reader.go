package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

//We fill in the byte of given length with 'A' and return the length and nil error
//The error returned is nil since we need to emit an infinite 'A'
//Typically we would have to return a EOF error
func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
