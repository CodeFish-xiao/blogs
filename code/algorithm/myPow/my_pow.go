package myPow

func myPow(x float64, n int) float64 {
	if 0 == n {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n > 0 {
		// 位运算，如果n是奇数，那么res乘以x
		if n&1 == 1 {
			res *= x
		}
		// x平方
		x *= x
		// 位运算，n右移一位，相当于n/2，因为n是整数，所以右移一位相当于n/2，这样循环下去，直到n=0
		n >>= 1
	}
	return res
}

// 递归
func myPow2(x float64, n int) float64 {
	if 0 == n {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n%2 == 0 {
		return myPow2(x*x, n/2)
	}
	return x * myPow2(x*x, n/2)
}
