package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fname := os.Args[1]
	length, err := fileLen(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Length of %s: %d", fname, length)
}

func fileLen(fname string) (int, error) {
	f, err := os.Open(fname)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	var (
		bytes int
		b     int
	)
	for {
		b, err = f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		bytes += b
	}
	return bytes, nil
}
