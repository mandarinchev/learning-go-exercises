// Dup2 prints the text of each line that appears more than
// once in the standard input or a set of files, preceded by its count,
// and followed by the name of each file where it occurred
package main

import (
	"bufio"
	"fmt"
	"os"
)

type UsageDescriptor struct {
	uses  int
	files map[string]bool
}

func main() {
	counts := make(map[string]*UsageDescriptor)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
				f.Close()
			}
		}
	}
	for line, u := range counts {
		if u.uses > 1 {
			fmt.Printf("%d\t%s\t", u.uses, line)
			for fileName, _ := range u.files {
				fmt.Print(fileName + " ")
			}
			fmt.Println()
		}
	}
}
func countLines(f *os.File, counts map[string]*UsageDescriptor) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		currentLine := input.Text()
		if counts[currentLine] == nil {
			var ud UsageDescriptor
			counts[currentLine] = &ud
			counts[currentLine].files = make(map[string]bool)
		}
		counts[currentLine].uses++
		if counts[currentLine].files[f.Name()] == false {
			counts[currentLine].files[f.Name()] = true
		}
	}
}
