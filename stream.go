package spritz

import "crypto/cipher"

// NewStream returns a new instance of the Spritz cipher using the given key.
func NewStream(key []byte) cipher.Stream {
	var s state
	s.initialize(256)

	// key setup
	s.absorbBytes(key)
	if s.a > 0 {
		s.shuffle()
	}

	return stream{s: &s}
}

type stream struct {
	s *state
}

func (s stream) XORKeyStream(dst, src []byte) {
	for i, v := range src {
		dst[i] = v ^ byte(s.s.drip())
	}
}
