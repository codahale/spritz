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
