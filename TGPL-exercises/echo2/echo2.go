// Echo2 prints its command-line arguments and their indexes line-by-line
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%v: %v\n", i, arg)
	}
}
