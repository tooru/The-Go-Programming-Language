package main

import (
    "bufio"
    "crypto/sha256"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    buf := make([]byte, 1024)
    digest := sha256.New()

    for {
        len, err := reader.Read(buf)
        if len == 0 {
            break;
        }
        if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			os.Exit(1)
        }
        digest.Write(buf[:len])
    }
    fmt.Printf("%x\n", digest.Sum(nil))
}
