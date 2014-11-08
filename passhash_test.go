package spritz_test

import (
	"testing"

	"github.com/codahale/spritz"
)

func BenchmarkPasswordHash15(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 15, 32)
	}
}

func BenchmarkPasswordHash16(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 16, 32)
	}
}

func BenchmarkPasswordHash17(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 17, 32)
	}
}

func BenchmarkPasswordHash18(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 18, 32)
	}
}

func BenchmarkPasswordHash19(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 19, 32)
	}
}

func BenchmarkPasswordHash20(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 20, 32)
	}
}
