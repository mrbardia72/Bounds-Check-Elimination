package main

func f1(s []int) {
	_ = s[0] // line 5: bounds check
	_ = s[1] // line 6: bounds check
	_ = s[2] // line 7: bounds check
}

func f2(s []int) {
	_ = s[2] // line 11: bounds check
	_ = s[1] // line 12: bounds check eliminated!
	_ = s[0] // line 13: bounds check eliminated!
}

func f3(s []int, index int) {
	_ = s[index] // line 17: bounds check
	_ = s[index] // line 18: bounds check eliminated!
}

func f5(s []int, k int) {
	_ = s[k] // line 17: bounds check
	_ = s[k] // line 18: bounds check eliminated!
	_ = s[k] // line 18: bounds check eliminated!
	_ = s[k] // line 18: bounds check eliminated!
}

func f4(a [5]int) {
	_ = a[4] // line 22: bounds check eliminated!
}

func main() {}
