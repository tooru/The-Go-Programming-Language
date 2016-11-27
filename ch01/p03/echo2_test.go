package ch01_03
//package main

import (
    "testing"
    "strconv"
    "strings"
)

func BenchmarkJoinStrings(b *testing.B) {
    var ss []string
    for i := 0; i < 100; i++ {
        ss = append(ss, strconv.Itoa(i))
    }
    
    for i := 0; i < b.N; i++ {
        JoinStrings(ss)
    }
}

func BenchmarkJoin(b *testing.B) {
    var ss []string
    for i := 0; i < 100; i++ {
        ss = append(ss, strconv.Itoa(i))
    }
    
    for i := 0; i < b.N; i++ {
        strings.Join(ss, " ")
    }
}
