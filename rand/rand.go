package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const SessionTokeBytes = 32

/*
bytes generates a randon []byte
n is the number of bytes
*/
func bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)

	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}

	if nRead < n {
		return nil, fmt.Errorf("bytes: didn't read enough random bytes")
	}

	return b, nil
}

/*
Strings returns a random string using crypt/rand
n is the number of bytes being used to generate the random string
*/
func strings(n int) (string, error) {
	b, err := bytes(n)
	if err != nil {
		return "", fmt.Errorf("strings: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func SessionToken() (string, error) {
	return strings(SessionTokeBytes)
}
