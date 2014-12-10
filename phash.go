/*
Package phash implements simple functions for saltling, hashing and later
verifying passwords against hashes. It is intended to be complatible
with the Node.js package "password-hash".
*/
package phash

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"hash"
	"strconv"
	"strings"
)

var stringToHash = map[string]crypto.Hash{
	"sha384":    crypto.SHA384,
	"md5sha1":   crypto.MD5SHA1,
	"sha512":    crypto.SHA512,
	"sha256":    crypto.SHA256,
	"md5":       crypto.MD5,
	"md4":       crypto.MD4,
	"sha1":      crypto.SHA1,
	"sha224":    crypto.SHA224,
	"ripemd160": crypto.RIPEMD160,
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generateSalt(length uint) []byte {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return bytes
}

func generateHash(h func() hash.Hash, salt []byte, password string, iter uint) string {
	hash := password
	for i := uint(0); i < iter; i++ {
		hasher := hmac.New(h, salt)
		hasher.Write([]byte(hash))
		hash = fmt.Sprintf("%x", hasher.Sum(nil))
	}
	return hash
}

// Generate converts a plain text password to a hash with a salt. It allows the
// caller to specify the algorithm, salt length, and number of times to apply
// the algorithm.
func Generate(password, algorithm string, saltLen, iter uint) (string, error) {
	a, ok := stringToHash[algorithm]
	if !ok {
		return "", fmt.Errorf("Unrecognized hash")
	}
	if !a.Available() {
		return "", fmt.Errorf("Hash algorithm not available: %s", algorithm)
	}
	salt := generateSalt(saltLen)
	hash := generateHash(a.New, salt, password, iter)
	i := strconv.FormatUint(uint64(iter), 10)
	return algorithm + "$" + string(salt) + "$" + i + "$" + hash, nil
}

// Gen converts a plain text password to a hash with a salt. It uses sha1, a
// salt lenght of 8, and applies the algorithm once.
func Gen(password string) string {
	salt := generateSalt(8)
	hash := generateHash(sha1.New, salt, password, 1)
	return "sha1$" + string(salt) + "$1$" + hash
}

// Verify compares a plain text password against a hash with a salt and returns
// true if the two match. If the hash algorithm used for the hash isn't linked
// to the binary, Verify automatically returns false. See
// http://golang.org/pkg/crypto/#RegisterHash for more details.
func Verify(password, hashedPassword string) bool {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 4 {
		return false
	}
	iter, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return false
	}
	h, ok := stringToHash[strings.ToLower(parts[0])]
	if !ok || !h.Available() {
		return false
	}
	hashCompare := generateHash(h.New, []byte(parts[1]), password, uint(iter))
	if hashCompare == parts[3] {
		return true
	}
	return false
}
