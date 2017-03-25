package p07_reverseUTF8Bytes

import (
    "testing"
)

func assertEquals(t *testing.T, x, y []byte) {
    if len(x) != len(y) {
        t.Errorf("%v %v", x, y)
        return
    }
    for i := 0; i < len(x); i++ {
        if x[i] != y[i] {
            t.Errorf("%v %v", x, y)
            return
        }
    }
}

func TestUniqSapce(t *testing.T) {
    assertEquals(t, nil, ReverseUTF8Bytes(nil))

    assertEquals(t, []byte("a"), ReverseUTF8Bytes([]byte("a")))
    assertEquals(t, []byte("あ"), ReverseUTF8Bytes([]byte("あ")))

    assertEquals(t, []byte("あa"), ReverseUTF8Bytes([]byte("aあ")))

    assertEquals(t, []byte("お0え987う654い32あ1"), ReverseUTF8Bytes([]byte("1あ23い456う789え0お")))
}
