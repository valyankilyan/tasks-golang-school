package sumdecimal

import (
	"math"
	"math/big"
	"strconv"
)

func SumDecimal(c int) int {
	if c <= 1 {
		return 0
	}
	precision := 1000
	ans_int := strconv.Itoa(int(math.Sqrt(float64(c))))

	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)+3), nil)
	a := big.NewInt(5 * int64(c))
	b := big.NewInt(5)
	five := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)

	for b.Cmp(limit) < 0 {
		if a.Cmp(b) < 0 {
			a.Mul(a, hundred)
			tmp := new(big.Int).Div(b, ten)
			tmp.Mul(tmp, hundred)
			b.Add(tmp, five)
		} else {
			a.Sub(a, b)
			b.Add(b, ten)
		}
	}
	b.Div(b, hundred)

	str := b.String()[len(ans_int):]
	ans := 0

	for i := 0; i < len(str) && i < precision; i++ {
		ans += int(byte(str[i] - byte('0')))
	}

	return ans
}
