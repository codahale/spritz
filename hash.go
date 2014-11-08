package spritz

import "hash"

// NewHash returns a new instance of the Spritz hash with the given output size.
func NewHash(size int) hash.Hash {
	var s state
	h := hasher{size: size, s: &s}
	h.Reset()
	return h
}

type hasher struct {
	size int
	s    *state
}

func (h hasher) Sum(b []byte) []byte {
	s := *h.s // make a local copy
	s.absorbStop()
	s.absorbValue(h.size)

	out := make([]byte, h.size)
	s.squeeze(out)

	return append(b, out...)
}

func (h hasher) Write(p []byte) (int, error) {
	h.s.absorb(p)
	return len(p), nil
}

func (h hasher) Size() int {
	return h.size
}

func (h hasher) Reset() {
	h.s.initialize(256)
}

func (hasher) BlockSize() int {
	return 1 // single byte
}
