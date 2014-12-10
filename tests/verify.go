package main

import (
	"fmt"
	"os"

	"github.com/yhat/phash"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "code.google.com/p/go.crypto/md4"
	_ "code.google.com/p/go.crypto/ripemd160"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run verify.go <password> <hash>")
		os.Exit(2)
	}
	password := os.Args[1]
	hash := os.Args[2]
	if !phash.Verify(password, hash) {
		fmt.Fprintf(os.Stderr, "%s did not match %s\n", password, hash)
		os.Exit(2)
	}
}
