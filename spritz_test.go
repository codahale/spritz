package spritz_test

import (
	"bytes"
	"testing"

	"github.com/codahale/spritz"
)

func TestStream(t *testing.T) {
	v := []byte{'a', 'r', 'c', 'f', 'o', 'u', 'r'}
	s := spritz.NewStream(v)

	out := make([]byte, 32)
	s.XORKeyStream(out, out)

	expected := []byte{
		0x1a, 0xfa, 0x8b, 0x5e, 0xe3, 0x37, 0xdb, 0xc7, 0x22, 0x59, 0x7f, 0x0f,
		0xdc, 0x3a, 0x42, 0xc7, 0x75, 0x4b, 0xf1, 0x03, 0x6f, 0x54, 0xfb, 0x4a,
		0xeb, 0x03, 0x35, 0xd4, 0xa4, 0xe9, 0xa3, 0x6e,
	}

	if !bytes.Equal(out, expected) {
		t.Errorf("Was \n%x\n but expected\n%x", out, expected)
	}
}

func TestHash(t *testing.T) {
	v := []byte{'a', 'r', 'c', 'f', 'o', 'u', 'r'}
	h := spritz.NewHash(32)
	_, _ = h.Write(v)
	out := h.Sum(nil)

	expected := []byte{
		0xff, 0x8c, 0xf2, 0x68, 0x9, 0x4c, 0x87, 0xb9, 0x5f, 0x74, 0xce, 0x6f,
		0xee, 0x9d, 0x30, 0x3, 0xa5, 0xf9, 0xfe, 0x69, 0x44, 0x65, 0x3c, 0xd5,
		0xe, 0x66, 0xbf, 0x18, 0x9c, 0x63, 0xf6, 0x99,
	}

	if !bytes.Equal(out, expected) {
		t.Errorf("Was \n%x\n but expected\n%x", out, expected)
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
