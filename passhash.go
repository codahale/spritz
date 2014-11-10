package spritz

// InsecurePasswordHash calculates a CPU- and memory-hard hash of the given
// password and salt. It takes two exponential parameters, M and T, which
// determine the memory and CPU costs. It also takes the length of the hash in
// bytes.
//
// N.B.: THIS IS A TOTALLY EXPERIMENTAL ALGORITHM WHICH I WROTE BEFORE I'D HAD
// ANY COFFEE. DO NOT USE HACKY ALGORITHMS DESIGNED BY UNCAFFEINATED
// NON-CRYPTOGRAPHERS.
func InsecurePasswordHash(password, salt []byte, m, t, n uint) []byte {
	// initialize to 256*(2**m) bytes
	var s state
	s.initialize(1 << (m + 8))

	// absorb the password
	s.absorb(password)

	// absorb the salt
	s.absorbStop()
	s.absorb(salt)

	// absorb the length
	s.absorbStop()
	s.absorbByte(int(n))

	// iterate through 2**t bytes of output
	for i := 0; i < 1<<uint(t); i++ {
		s.drip()
	}

	// squeeze out the digest
	out := make([]byte, n)
	s.squeeze(out)
	return out
}
