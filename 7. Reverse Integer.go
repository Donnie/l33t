func reverse(x int) int {
	mul, rev := 1, 0
	if x < 0 {
		mul = -1
	}
	unsigned := x * mul

	for unsigned > 0 {
		carry := unsigned % 10
		rev = rev*10 + carry
		unsigned = unsigned / 10
	}

	if rev > 2147483648 {
		return 0
	}
	return rev * mul
}
