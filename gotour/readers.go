package main

import (
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mR MyReader) Read(byteArray []byte) (int, error) {
	strr := strings.NewReader("")

	aByte := make([]byte, 'A')

	for {
		aReader, ReadError := strr.Read(aByte)

		if ReadError == io.EOF {
			return aReader, ReadError
		}
	}
}

func main() {
	reader.Validate(MyReader{})
}
