package romannumerals

import "fmt"

// Показалось более правильным, чем мап, потому что размер пмассива всего 7
// поэтому поиск будет осуществлятся за Θ(n)
var numerals = [...]struct {
	decimal int
	roman   rune
}{
	{1, 'I'},
	{5, 'V'},
	{10, 'X'},
	{50, 'L'},
	{100, 'C'},
	{500, 'D'},
	{1000, 'M'},
}

func GetDecimal(roman rune) (int, error) {
	for _, num := range numerals {
		if num.roman == roman {
			return num.decimal, nil
		}
	}
	return 0, fmt.Errorf("can't find decimal")
}

func BiggestIndex(roman []rune) (int, error) {
	if len(roman) == 0 {
		return 0, fmt.Errorf("can't find biggest index in empty string")
	}
	mx, ind := 0, 0
	for i, r := range roman {
		if dec, err := GetDecimal(r); dec > mx {
			if err != nil {
				return 0, err
			}
			mx = dec
			ind = i
		}
	}
	return ind, nil
}

func DecodeSlice(roman []rune) (int, error) {
	if len(roman) == 0 {
		return 0, nil
	} else {
		biggest, err := BiggestIndex(roman)
		if err != nil {
			return 0, err
		}

		left, err := DecodeSlice(roman[:biggest])
		if err != nil {
			return 0, err
		}

		right, err := DecodeSlice(roman[biggest+1:])
		if err != nil {
			return 0, err
		}

		centre, err := GetDecimal(roman[biggest])
		if err != nil {
			return 0, err
		}

		dec := centre - left + right

		return dec, nil
	}
}

func Decode(s string) (int, bool) {
	// написал вот этот костыль, потому что это не вписывается в мою концепцию
	// пустая строка должна возвращать ошибку, хотя по идее это 0
	if len(s) == 0 {
		return 0, false
	}
	ans, err := DecodeSlice([]rune(s))
	return ans, err == nil
}

func FindMaxIndexRoman(decimal int) (int, error) {
	ind := 0
	for i, num := range numerals {
		if num.decimal <= decimal {
			ind = i
		}
	}

	return ind, nil
}

func EncodeSlice(decimal int) ([]rune, error) {
	if decimal == 0 {
		return []rune{}, nil
	}

	mxind, err := FindMaxIndexRoman(decimal)
	if err != nil {
		return []rune{}, err
	}

	mxlen := decimal / numerals[mxind].decimal
	if mxlen == 4 && mxind != len(numerals)-1 {
		ap, err := EncodeSlice(decimal + numerals[mxind].decimal)
		if err != nil {
			return []rune{}, err
		}
		return append([]rune{numerals[mxind].roman}, ap...), nil
	} else {
		ap, err := EncodeSlice(decimal % numerals[mxind].decimal)
		if err != nil {
			return []rune{}, err
		}

		cur := make([]rune, mxlen)
		for i := 0; i < mxlen; i++ {
			cur[i] = numerals[mxind].roman
		}

		return append(cur, ap...), nil
	}
}

func Encode(n int) (string, bool) {
	if n <= 0 {
		return "", false
	}

	ans, err := EncodeSlice(n)
	if err != nil {
		return "", false
	}
	return string(ans), true
}
