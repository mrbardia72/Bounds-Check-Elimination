package main

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
