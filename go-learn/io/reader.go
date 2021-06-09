package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (n int, err error) {

	for ind, _ := range b {
		b[ind] = 'A'
	}

	n = len(b)

	return
}

func main() {
	reader.Validate(MyReader{})
}
