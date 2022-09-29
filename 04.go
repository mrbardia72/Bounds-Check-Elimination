package main

//Slice accesses offset from a base index 1
func bce4(b []byte, n int) {
	b[n+0] = byte(1)
	b[n+1] = byte(1 >> 8)
	b[n+2] = byte(1 >> 16)
	b[n+3] = byte(1 >> 24)
	b[n+4] = byte(1 >> 32)
	b[n+5] = byte(1 >> 40)
}

//Slice accesses offset from a base index 2
func bce5(b []byte, n int) {
	_ = b[n+5]
	b[n+0] = byte(1)
	b[n+1] = byte(1 >> 8)
	b[n+2] = byte(1 >> 16)
	b[n+3] = byte(1 >> 24)
	b[n+4] = byte(1 >> 32)
	b[n+5] = byte(1 >> 40)
}

//Slice accesses offset from a base index 2
func bce6(b []byte, n int) {
	b = b[n : n+6]
	b[0] = byte(1)
	b[1] = byte(1 >> 8)
	b[2] = byte(1 >> 16)
	b[3] = byte(1 >> 24)
	b[4] = byte(1 >> 32)
	b[5] = byte(1 >> 40)
}
