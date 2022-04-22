package main

import (
	"fmt"
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	var a, c = 4, 4
	for i := 0; i < b.N; i++ {
		Encrypt(fmt.Sprintf("Test text %d", i), a, c)
	}
}

