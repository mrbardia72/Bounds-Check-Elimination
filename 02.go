package main

//Loop upper bound is known
func bce1(b []byte, n int) {
	_ = b[n-1]
	for i := 0; i < n; i++ {
		b[i] = 9
	}
}

//Contiguous slice access 1
func bce2(b []byte, v uint32) {
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}

//Contiguous slice access 2
func bce3(b []byte, v uint32) {
	_ = b[3]
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}

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

//Iterate multiple slices with known upper bound 1
func bce7(a, b, c []int) {
	for i := range a {
		a[i] = b[i] + c[i]
	}
}

//Iterate multiple slices with known upper bound 2
func bce8(a, b, c []int) {
	_ = b[len(a)-1]
	_ = c[len(a)-1]
	for i := 0; i < len(a); i++ {
		for i := range a {
			a[i] = b[i] + c[i]
		}
	}
}

//Slice range can be derived 1
func bce9(b []byte, h []int32) {
	for _, t := range b {
		h[t]++
	}
}

//Slice range can be derived
func bce10(b []byte, h []int32) {
	h = h[:256]
	for _, t := range b {
		h[t]++
	}
}
