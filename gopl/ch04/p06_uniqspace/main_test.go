package p06_uniqspace

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
    assertEquals(t, nil, UniqSpace(nil))

    assertEquals(t, []byte("a"), UniqSpace([]byte("a")))

    assertEquals(t, []byte(" "), UniqSpace([]byte(" ")))
    assertEquals(t, []byte(" "), UniqSpace([]byte("  ")))
    assertEquals(t, []byte(" "), UniqSpace([]byte(" \t")))
    assertEquals(t, []byte(" "), UniqSpace([]byte(" \t\n")))
    assertEquals(t, []byte(" "), UniqSpace([]byte(" \t\n\v\f\r\u0085\u00a0")))

    assertEquals(t, []byte(" a b "), UniqSpace([]byte(" \t\n\v\f\r\u0085\u00a0a       b\t\n")))
}
