package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkEncrypt(b *testing.B) {
	var a, c = 4, 4
	var dictionary = "abcdefghijklmnopqrstuvwxyz "
	st := NewAffineCipherManager(dictionary)
	for i := 0; i < b.N; i++ {
		st.Encrypt(fmt.Sprintf("Test text %d", i), a, c)
	}
}

func TestDecrypt(t *testing.T) {
	var a, b = 4, 4
	var dictionary = "abcdefghijklmnopqrstuvwxyz "
	var cases = []struct{
		encrypted string
		decrypted string
	}{
		{
			encrypted: "fuvvgalgsvq",
			decrypted: "hello world",
		},
		{
			encrypted: "wusubearsewehe",
			decrypted: "serega krasava",
		},
	}

	st := NewAffineCipherManager(dictionary)
	for _, c := range cases {
		decrypted := st.Decrypt(c.encrypted, a, b)
		assert.Equal(t, c.decrypted, decrypted)
	}
}

