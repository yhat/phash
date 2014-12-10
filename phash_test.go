package phash

import (
	"crypto/rand"
	"testing"
)

func randPasswd(n uint) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func TestGen(t *testing.T) {
	for i := 1; i < 1000; i++ {
		pass, err := randPasswd(uint(i))
		if err != nil {
			t.Fatal(err)
		}
		hash := Gen(pass)
		if !Verify(pass, hash) {
			t.Errorf("Exepcted %s to pass", pass)
			return
		}
	}
}
