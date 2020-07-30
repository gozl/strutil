package strutil

import (
	"math/rand"
	srand "crypto/rand"
	"time"
	"unsafe"
	"bytes"
	"encoding/binary"
)

const (
	alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	secretCharset = "abcdefghinpqrtABCDEFGHJKLMNPRSTUWXY1234567890~@_/+:"
	letterIdxBits = 6                      // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits     // #
)

// RandomBase64 returns a random string of the specified length in the character range 0-9a-zA-Z+/
func RandomBase64(n int) string {
	return Random(n, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
}

// RandomBase62 returns a random string of the specified length in the character range 0-9a-zA-Z
func RandomBase62(n int) string {
	return Random(n, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// RandomBase36 returns a random string of the specified length in the character range 0-9a-z
func RandomBase36(n int) string {
	return Random(n, "0123456789abcdefghijklmnopqrstuvwxyz")
}

// RandomHex returns a random string of the specified length in the character range 0-9a-f
func RandomHex(n int) string {
	return Random(n, "0123456789abcdef")
}

// Random returns a random string of the specified length. Uses alphanumeric characters by default.
func Random(n int, charset ...string) string {
	b := make([]byte, n)

	var letterRunes []rune
	if len(charset) == 0 {
		letterRunes = []rune(alphanumeric)
	} else {
		letterRunes = []rune(charset[0])
	}

	// a randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	randSrc := rand.NewSource(time.Now().UnixNano())
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterRunes) {
			b[i] = byte(letterRunes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// Secret generates a secure random string using crypto/rand. Possible characters 
// are `abcdefghinpqrt`, `ABCDEFGHJKLMNPRSTUWXY`, 0-9, and the symbols `~@_/+:`. Password policy 
// compliance is not handled by this function. Error may be returned due to crypto/rand Read().
func Secret(n int, charset ...string) (string, error) {
	b := make([]byte, n * 4)
	_, err := srand.Read(b)
	if err != nil {
		return "", err
	}

	var letterRunes []rune
	if len(charset) == 0 {
		letterRunes = []rune(secretCharset)
	} else {
		letterRunes = []rune(charset[0])
	}

	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))

	for i := 0; i < (n * 4); i+=4 {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(b[i:i+4]) % l])
	}
	return bb.String(), nil
}
