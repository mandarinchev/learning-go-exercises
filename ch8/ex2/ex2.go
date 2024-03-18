package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type IntStringer int

func (is IntStringer) String() string {
	return strconv.Itoa(int(is))
}

type FloatStringer float64

func (fs FloatStringer) String() string {
	return fmt.Sprintf("%f", fs)
}

func PrintPrintable[T Printable](p T) {
	fmt.Println(p)
}

func main() {
	var i IntStringer = 10
	PrintPrintable(i)

	var f FloatStringer = 5.2
	PrintPrintable(f)
}
