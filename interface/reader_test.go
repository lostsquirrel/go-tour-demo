package iface_demo

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReader2(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}
