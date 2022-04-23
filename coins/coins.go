package coins

func Piles(n int) int {
	return (1 << (n - 1)) - 1
}
