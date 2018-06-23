package main

import (
	"io"
	"strings"
			"os"
)

var rot13 = map[byte]byte {
'A': 'N',
'B': 'O',
'C': 'P',
'D': 'Q',
'E': 'R',
'F': 'S',
'G': 'T',
'H': 'U',
'I': 'V',
'J': 'W',
'K': 'X',
'L': 'Y',
'M': 'Z',
'a': 'n',
'b': 'o',
'c': 'p',
'd': 'q',
'e': 'r',
'f': 's',
'g': 't',
'h': 'u',
'i': 'v',
'j': 'w',
'k': 'x',
'l': 'y',
'm': 'z',
'N': 'A',
'O': 'B',
'P': 'C',
'Q': 'D',
'R': 'E',
'S': 'F',
'T': 'G',
'U': 'H',
'V': 'I',
'W': 'J',
'X': 'K',
'Y': 'L',
'Z': 'M',
'n': 'a',
'o': 'b',
'p': 'c',
'q': 'd',
'r': 'e',
's': 'f',
't': 'g',
'u': 'h',
'v': 'i',
'w': 'j',
'x': 'k',
'y': 'l',
'z': 'm',
}
type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i := 0; i < n; i++ {
		bx, ok := rot13[b[i]]
		if ok {
			b[i] = bx
		}

	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)

}

