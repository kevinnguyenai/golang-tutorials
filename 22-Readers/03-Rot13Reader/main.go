package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotr rot13Reader) Read(b []byte) (int, error) {
	n, err := rotr.r.Read(b)
	for i:=0;i<n;i++ {
		switch {
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = 'a'+(b[i]-'a'+13)%26
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = 'A'+(b[i]-'a'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
