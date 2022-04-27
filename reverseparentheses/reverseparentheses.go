package reverseparentheses

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() stack {
	return s[:len(s)-1]
}

func (s stack) Top() int {
	return s[len(s)-1]
}

func is_brackets(r byte) bool {
	return r == '(' || r == ')'
}

func reverse(chars []byte, from, to int) []byte {
	l, r := from, to
	for l < r {
		for is_brackets(chars[l]) {
			l++
		}
		for is_brackets(chars[r]) {
			r--
		}
		if l < r {
			temp := chars[l]
			chars[l] = chars[r]
			chars[r] = temp
			l++
			r--
		}
	}

	return chars
}

func Reverse(s string) string {
	st := make(stack, 0)
	chars := []byte(s)
	brackets_count := 0
	for i, c := range chars {
		if c == '(' {
			st = st.Push(i)
			brackets_count++
		} else if c == ')' {
			reverse(chars, st.Top()+1, i-1)
			st = st.Pop()
			brackets_count++
		}
	}

	out := make([]byte, len(s)-brackets_count)
	out_index := 0
	for _, c := range chars {
		if !is_brackets(c) {
			out[out_index] = c
			out_index++
		}
	}

	return string(out)
}
