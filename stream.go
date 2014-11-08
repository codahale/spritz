package spritz

import "crypto/cipher"

// NewStream returns a new instance of the Spritz cipher using the given key.
func NewStream(key []byte) cipher.Stream {
	return NewStreamWithIV(key, nil)
}

// NewStreamWithIV returns a new instance of the Spritz cipher using the given
// key and initialization vector.
func NewStreamWithIV(key, iv []byte) cipher.Stream {
	var s state
	s.initialize(256)

	// key setup
	s.absorb(key)
	if s.a > 0 {
		s.shuffle()
	}
	if iv != nil {
		s.absorbStop()
		s.absorb(iv)
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
