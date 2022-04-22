package iface_demo

import (
	"testing"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

//  Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for i < len(b) {
		b[i] = 'A'
		i++
	}
	return i, nil
}

func TestReader(t *testing.T) {
	reader.Validate(MyReader{})
}
