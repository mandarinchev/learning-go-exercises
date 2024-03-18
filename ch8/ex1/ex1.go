package main

import "fmt"

type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

func doubleVal[T Numeric](n T) T {
    return n * 2
}

func main() {
    fmt.Println(doubleVal(25))
    fmt.Println(doubleVal(2.6))
    fmt.Println(doubleVal(100000000000000000))
}
