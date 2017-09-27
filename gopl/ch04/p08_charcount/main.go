// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	//	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

func CountChar(rd io.Reader) (map[string]int, error) {
	counts := make(map[string]int) // counts of Unicode characters

	in := bufio.NewReader(rd)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			return nil, err
		}
		switch {
		case unicode.IsControl(r):
			counts["control"]++
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsMark(r):
			counts["mark"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsSymbol(r):
			counts["symbol"]++
		default:
			fmt.Printf("%v\n", string(r))
			counts["other"]++
			//return nil, errors.New("unknown type: " + string(r))
		}
	}
	return counts, nil
}

func main() {
	counts, err := CountChar(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

//!-
