package spritz

// InsecurePasswordHash calculates a CPU- and memory-hard hash of the given
// password and salt. It takes an exponential parameter, m, which determines
// both CPU and memory cost. It also takes the length of the hash in bytes.
//
// N.B.: THIS IS A TOTALLY EXPERIMENTAL ALGORITHM WHICH I WROTE BEFORE I'D HAD
// ANY COFFEE. DO NOT USE HACKY ALGORITHMS DESIGNED BY UNCAFFEINATED
// NON-CRYPTOGRAPHERS.
func InsecurePasswordHash(password, salt []byte, m, n uint) []byte {
	// initialize to 2**m bytes
	var s state
	s.initialize(1 << m)

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
	out := make([]byte, int(n))
	s.squeeze(out)
	return out
}
