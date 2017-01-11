// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"

	"gopl/ch02/p02_unitconv"
)

func main() {
    files := os.Args[1:]

    if len(files) == 0 {
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            text := input.Text()
            printUnit(text)
        }
    } else {
        for _, arg := range files {
            printUnit(arg)
        }
	}
}

func printUnit(s string) {
    t, err := strconv.ParseFloat(s, 64)
    if err != nil {
        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
        os.Exit(1)
    }
    c  := p02_unitconv.Celsius(t)
    fh := p02_unitconv.Fahrenheit(t)
    m  := p02_unitconv.Meter(t)
    fe := p02_unitconv.Feet(t)
    kg := p02_unitconv.Kirogram(t)
    p  := p02_unitconv.Pond(t)
    fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
        c, p02_unitconv.CToF(c), fh, p02_unitconv.FToC(fh),
        m, p02_unitconv.MToF(m), fe, p02_unitconv.FToM(fe),
        kg, p02_unitconv.KgToP(kg), p, p02_unitconv.PToKg(p),
    )
}    

//!-
