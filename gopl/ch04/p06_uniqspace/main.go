package p06_uniqspace

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func UniqSpace(b []byte) []byte {
	var buf bytes.Buffer

	var prevIsSpace = false

	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		runeIsSpace := unicode.IsSpace(r)
		if runeIsSpace {
			if prevIsSpace {
				prevIsSpace = runeIsSpace
				i += size
				continue
			}
			r = ' '
		}
		buf.WriteRune(r)
		prevIsSpace = runeIsSpace
		i += size
	}
	return buf.Bytes()
}
