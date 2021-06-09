package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (n int, err error) {

	length := len(b)
	bb := make([]byte, length)

	n, err = rot13.r.Read(bb)

	for ind, _ := range bb {
		if bb[ind] > 'M' && bb[ind] < 'Z' {
			b[ind] = bb[ind] - 13
		} else if bb[ind] > 'm' && bb[ind] < 'z' {
			b[ind] = bb[ind] - 13
		} else {
			b[ind] = bb[ind] + 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
