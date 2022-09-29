package main

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
