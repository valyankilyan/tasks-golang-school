package coins

func Rec(left, mx int) int {
	if left == 0 {
		return 1
	}
	ans := 0
	for i := mx; i <= left; i++ {
		ans += Rec(left-i, i)
	}
	return ans
}

func Piles(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans += Rec(n-i, i)
	}
	return ans
}
