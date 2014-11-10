package spritz_test

import (
	"testing"

	"github.com/codahale/spritz"
)

func BenchmarkPasswordHash(b *testing.B) {
	v := []byte("hello this is a password")
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		spritz.InsecurePasswordHash(v, nil, 10, 11, 32)
	}
}
