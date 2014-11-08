package spritz_test

import (
	"bytes"
	"testing"

	"github.com/codahale/spritz"
)

// test cases courtesy of github.com/dgryski/go-spritz

func TestStream(t *testing.T) {
	fixtures := []struct {
		key    string
		output []byte
	}{
		{"ABC", []byte{0x77, 0x9a, 0x8e, 0x01, 0xf9, 0xe9, 0xcb, 0xc0}},
		{"spam", []byte{0xf0, 0x60, 0x9a, 0x1d, 0xf1, 0x43, 0xce, 0xbf}},
		{"arcfour", []byte{0x1a, 0xfa, 0x8b, 0x5e, 0xe3, 0x37, 0xdb, 0xc7}},
	}

	for _, f := range fixtures {
		s := spritz.NewStream([]byte(f.key))

		out := make([]byte, len(f.output))
		s.XORKeyStream(out, out)

		if !bytes.Equal(out, f.output) {
			t.Errorf("Output for %q was \n%x\n but expected\n%x", f.key, out, f.output)
		}
	}
}

func TestHash(t *testing.T) {
	fixtures := []struct {
		key    string
		output []byte
	}{
		// PDF only provides first 8 bytes for a 32-byte hash
		{"ABC", []byte{0x02, 0x8f, 0xa2, 0xb4, 0x8b, 0x93, 0x4a, 0x18}},
		{"spam", []byte{0xac, 0xbb, 0xa0, 0x81, 0x3f, 0x30, 0x0d, 0x3a}},
		{"arcfour", []byte{0xff, 0x8c, 0xf2, 0x68, 0x09, 0x4c, 0x87, 0xb9}},
	}

	for _, f := range fixtures {
		h := spritz.NewHash(32)
		_, _ = h.Write([]byte(f.key))
		out := h.Sum(nil)[:len(f.output)]

		if !bytes.Equal(out, f.output) {
			t.Errorf("Output for %q was \n%x\n but expected\n%x", f.key, out, f.output)
		}
	}
}

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

func BenchmarkStream(b *testing.B) {
	v := []byte{'a', 'r', 'c', 'f', 'o', 'u', 'r'}
	s := spritz.NewStream(v)
	out := make([]byte, 1024)
	b.SetBytes(int64(len(out)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.XORKeyStream(out, out)
	}
}

func BenchmarkHash(b *testing.B) {
	h := spritz.NewHash(32)
	out := make([]byte, 1024)
	b.SetBytes(int64(len(out)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = h.Write(out)
	}
}
