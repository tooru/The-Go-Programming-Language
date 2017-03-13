package main

import (
    "bufio"
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "fmt"
    "hash"
    "os"
)

func main() {
    var length int

    flag.IntVar(&length, "l", 256, "specify the length of sha. 256, 384, 512")
    flag.Parse()

    var hash hash.Hash

    if length == 256 {
        hash = sha256.New()
    } else if length == 384 {
        hash = sha512.New384()
    } else if length == 512 {
        hash = sha512.New()
    } else {
        fmt.Fprintf(os.Stderr, "Unrecognized algorithm: %d\n", length)
        flag.PrintDefaults()
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)
    buf := make([]byte, 1024)

    for {
        len, err := reader.Read(buf)
        if len == 0 {
            break;
        }
        if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			os.Exit(1)
        }
        hash.Write(buf[:len])
    }
    fmt.Printf("%x\n", hash.Sum(nil))
}
