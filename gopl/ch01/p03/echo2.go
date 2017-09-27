package ch01_03

//package main

func JoinStrings(ss []string) {
	s, sep := "", ""
	for j := 0; j < len(ss); j++ {
		s += sep + ss[j]
		sep = " "
	}
}
