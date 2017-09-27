package p07_reverseUTF8Bytes

import (
	"bytes"
	"unicode/utf8"
)

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func ReverseUTF8Bytes(b []byte) []byte {
	var lens []int
	var bs = b

	for len(bs) > 0 {
		_, size := utf8.DecodeLastRune(bs)
		lens = append(lens, size)

		bs = bs[:len(bs)-size]
	}
	reverse(b)

	for i, j := 0, 0; i < len(lens); i++ {
		l := lens[i]
		reverse(b[j : j+l])
		j += l
	}

	return b
}

func ReverseUTF8Bytes2(b []byte) []byte {
	var buf bytes.Buffer

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		buf.WriteRune(r)

		b = b[:len(b)-size]
	}
	return buf.Bytes()
}
