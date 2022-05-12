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

func getDecimal(roman rune) (int, error) {
	for _, num := range numerals {
		if num.roman == roman {
			return num.decimal, nil
		}
	}
	return 0, fmt.Errorf("can't find decimal")
}

func biggestIndex(roman []rune) (int, error) {
	if len(roman) == 0 {
		return 0, fmt.Errorf("can't find biggest index in empty string")
	}
	mx, ind := 0, 0
	for i, r := range roman {
		if dec, err := getDecimal(r); dec > mx {
			if err != nil {
				return 0, err
			}
			mx = dec
			ind = i
		}
	}
	return ind, nil
}

func decodeSlice(roman []rune) (int, error) {
	if len(roman) == 0 {
		return 0, nil
	} else {
		biggest, err := biggestIndex(roman)
		if err != nil {
			return 0, err
		}

		left, err := decodeSlice(roman[:biggest])
		if err != nil {
			return 0, err
		}

		right, err := decodeSlice(roman[biggest+1:])
		if err != nil {
			return 0, err
		}

		centre, err := getDecimal(roman[biggest])
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
	ans, err := decodeSlice([]rune(s))
	return ans, err == nil
}

func findMaxIndexRoman(decimal int) (int, error) {
	ind := 0
	for i, num := range numerals {
		if num.decimal <= decimal {
			ind = i
		}
	}

	return ind, nil
}

func encodeSlice(decimal int) ([]rune, error) {
	if decimal == 0 {
		return []rune{}, nil
	}

	mxind, err := findMaxIndexRoman(decimal)
	if err != nil {
		return []rune{}, err
	}

	mxlen := decimal / numerals[mxind].decimal
	if mxlen == 4 && mxind != len(numerals)-1 {
		ap, err := encodeSlice(decimal + numerals[mxind].decimal)
		if err != nil {
			return []rune{}, err
		}
		return append([]rune{numerals[mxind].roman}, ap...), nil
	} else {
		ap, err := encodeSlice(decimal % numerals[mxind].decimal)
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

	ans, err := encodeSlice(n)
	if err != nil {
		return "", false
	}
	return string(ans), true
}
