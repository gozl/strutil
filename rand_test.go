package strutil

import (
	"testing"
	"fmt"
)

func TestRandomBase36(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := RandomBase36(6)
		fmt.Printf("random base36: %s\n", s)
	}
}


func TestRandomBase64(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := RandomBase64(6)
		fmt.Printf("random base64: %s\n", s)
	}
}

func TestRandomBase62(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := RandomBase62(6)
		fmt.Printf("random base62: %s\n", s)
	}
}

func TestRandomHex(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := RandomHex(6)
		fmt.Printf("random hex: %s\n", s)
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := Random(8, "abcdefghijABCD4321$#@")
		fmt.Printf("random string: %s\n", s)
	}

	for i := 0; i < 5; i++ {
		s := Random(8)
		fmt.Printf("random alphanum: %s\n", s)
	}
}

func TestSecret(t *testing.T) {
	for i := 0; i < 5; i++ {
		s, _ := Secret(8, "abcdefghijABCD4321$#@")
		fmt.Printf("secret custom: %s\n", s)
	}

	for i := 0; i < 5; i++ {
		s, _ := Secret(8)
		fmt.Printf("secret string: %s\n", s)
	}
}