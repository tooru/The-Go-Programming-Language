package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountWords(rd io.Reader) (map[string]int, error) {
	counts := make(map[string]int) // counts of Unicode characters

	in := bufio.NewScanner(rd)

	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		counts[word]++
	}

	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		return nil, err
	}
	return counts, nil
}

func main() {
	counts, err := CountWords(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "word: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
