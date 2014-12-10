package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yhat/phash"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "code.google.com/p/go.crypto/md4"
	_ "code.google.com/p/go.crypto/ripemd160"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run generate.go <password> <algorithm> <salt length> <iter>")
		os.Exit(2)
	}
	password := os.Args[1]
	algorithm := os.Args[2]
	saltLength, err := strconv.ParseUint(os.Args[3], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(4)
	}
	iter, err := strconv.ParseUint(os.Args[4], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(4)
	}
	hash, err := phash.Generate(password, algorithm, uint(saltLength), uint(iter))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(4)
	}
	fmt.Println(hash)
}
