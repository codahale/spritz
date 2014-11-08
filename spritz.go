// Package spritz provides a pure Go implementation of the Spritz stream cipher
// and hash.
//
// More details on the Spritz family of algorithms can be found here:
// http://people.csail.mit.edu/rivest/pubs/RS14.pdf.
package spritz

import "math"

type state struct {
	// these are all ints instead of bytes to allow for states > 256
	n, d             int // state size and nibble size
	s                []int
	a, i, j, k, w, z int
}

func (s *state) initialize(n int) {
	*s = state{
		s: make([]int, n),
		w: 1,
		n: n,
		d: int(math.Ceil(math.Sqrt(float64(n)))),
	}
	for i := range s.s {
		s.s[i] = i
	}
}

func (s *state) update() {
	s.i = (s.i + s.w) % s.n
	y := (s.j + s.s[s.i]) % s.n
	s.j = (s.k + s.s[y]) % s.n
	s.k = (s.i + s.k + s.s[s.j]) % s.n
	t := s.s[s.i]
	s.s[s.i] = s.s[s.j]
	s.s[s.j] = t
}

func (s *state) output() int {
	y1 := (s.z + s.k) % s.n
	x1 := (s.i + s.s[y1]) % s.n
	y2 := (s.j + s.s[x1]) % s.n
	s.z = s.s[y2]
	return s.z
}

func (s *state) crush() {
	for i := 0; i < s.n/2; i++ {
		y := (s.n - 1) - i
		x1 := s.s[i]
		x2 := s.s[y]
		if x1 > x2 {
			s.s[i] = x2
			s.s[y] = x1
		} else {
			s.s[i] = x1
			s.s[y] = x2
		}
	}
}

func (s *state) whip() {
	r := s.n * 2
	for i := 0; i < r; i++ {
		s.update()
	}
	s.w = (s.w + 2) % s.n
}

func (s *state) shuffle() {
	s.whip()
	s.crush()
	s.whip()
	s.crush()
	s.whip()
	s.a = 0
}

func (s *state) absorbStop() {
	if s.a == s.n/2 {
		s.shuffle()
	}
	s.a = (s.a + 1) % s.n
}

func (s *state) absorbNibble(x int) {
	if s.a == s.n/2 {
		s.shuffle()
	}
	y := (s.n/2 + x) % s.n
	t := s.s[s.a]
	s.s[s.a] = s.s[y]
	s.s[y] = t
	s.a = (s.a + 1) % s.n
}

func (s *state) absorbByte(b int) {
	s.absorbNibble(b % s.d) // LOW
	s.absorbNibble(b / s.d) // HIGH
}

func (s *state) absorb(msg []byte) {
	for _, v := range msg {
		s.absorbByte(int(v))
	}
}

func (s *state) drip() int {
	if s.a > 0 {
		s.shuffle()
	}
	s.update()
	return s.output()
}

func (s *state) squeeze(out []byte) {
	if s.a > 0 {
		s.shuffle()
	}
	for i := range out {
		out[i] = byte(s.drip())
	}
}
