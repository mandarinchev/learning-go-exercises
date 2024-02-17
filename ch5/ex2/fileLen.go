package main

import (
    "errors"
    "fmt"
    "io"
    "os"
)

func main() {
    fileName := os.Args[1]
    length, err := fileLen(fileName)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("Length of %s: %d", fileName, length)
}

func fileLen(fileName string) (int, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    buf := make([]byte, 1024)
    var (
        bytes int
        b     int
    )
    for {
        b, err = file.Read(buf)
        bytes += b
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            }
            return 0, err
        }
    }
    return bytes, nil
}
