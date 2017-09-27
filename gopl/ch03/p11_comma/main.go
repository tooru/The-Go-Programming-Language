// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`([-+])?(\d*)(\.\d+)?`)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(str string) string {
	result := re.FindStringSubmatch(str)

	match := len(result[0]) > 0
	s := result[1]
	d := result[2]
	f := result[3]

	if !match {
		return str
	}

	var buf bytes.Buffer

	buf.WriteString(s)

	n := len(d)
	i := n % 3
	c := ","
	if n == 3 {
		c = ""
	}

	buf.WriteString(d[:i])

	for ; i < n; i += 3 {
		buf.WriteString(c)
		buf.WriteString(d[i : i+3])
	}
	buf.WriteString(f)
	return buf.String()
}

//!-
