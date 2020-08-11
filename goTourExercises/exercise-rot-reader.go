package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt *rot13Reader) Read(b []byte) (int, error) {

	len, err := rt.r.Read(b)
	for i := 0; i < len; i++ {
		if b[i] >= 65 && b[i] < 91 {
			b[i] = 65 + ((b[i] - 65 + 13) % 26)
		} else if b[i] >= 97 && b[i] < 123 {
			b[i] = 97 + ((b[i] - 97 + 13) % 26)
			continue
		} else {
			continue
		}
	}

	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
