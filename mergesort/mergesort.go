package mergesort

func Merge(input []int, start, mid, end int) {
	var s1, s2 []int
	s1 = append(s1, input[start:mid]...)
	s2 = append(s2, input[mid:end]...)
	f1, f2 := 0, 0
	i := start
	for ; f1 < len(s1) && f2 < len(s2); i++ {
		if s1[f1] <= s2[f2] {
			input[i] = s1[f1]
			f1++
		} else {
			input[i] = s2[f2]
			f2++
		}
	}
	if f1 < len(s1) {
		for ; f1 < len(s1); f1++ {
			input[i] = s1[f1]
			i++
		}
	} else if f2 < len(s2) {
		for ; f2 < len(s2); f2++ {
			input[i] = s2[f2]
			i++
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	for delta := 1; delta < len(input); delta = delta << 1 {
		for i := 0; i < len(input) && i+delta < len(input); i += delta * 2 {
			Merge(input, i, i+delta, min(i+delta*2, len(input)))
		}
	}
	return input
}
