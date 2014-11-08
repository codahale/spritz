package spritz

// InsecurePasswordHash calculates a CPU- and memory-hard hash of the given
// password and salt. It takes a linear parameter, m, which determines both CPU
// and memory cost. It also takes the length of the hash in bytes.
//
// N.B.: THIS IS A TOTALLY EXPERIMENTAL ALGORITHM WHICH I WROTE BEFORE I'D HAD
// ANY COFFEE. DO NOT USE HACKY ALGORITHMS DESIGNED BY UNCAFFEINATED
// NON-CRYPTOGRAPHERS.
func InsecurePasswordHash(password, salt []byte, m, n int) []byte {
	// initialize to 256*m bytes
	var s state
	s.initialize(256 * m)

	// absorb the password
	s.absorb(password)
	if s.a > 0 {
		s.shuffle()
	}
	s.absorbStop()

	// absorb the salt
	s.absorb(salt)
	if s.a > 0 {
		s.shuffle()
	}
	s.absorbStop()

	// absorb the length
	s.absorbByte(int(n))
	s.absorbStop()

	// squeeze out the digest
	out := make([]byte, n)
	s.squeeze(out)
	return out
}
