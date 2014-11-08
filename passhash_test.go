package spritz_test

import (
	"testing"

	"github.com/codahale/spritz"
)

func BenchmarkPasswordHash256(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 256, 32)
	}
}

func BenchmarkPasswordHash512(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 512, 32)
	}
}

func BenchmarkPasswordHash1024(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 1024, 32)
	}
}
