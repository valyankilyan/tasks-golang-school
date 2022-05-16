package secretmessage

import "sort"

type LetterFrequency struct {
	letter    rune
	frequency int
}

// Decode func
func Decode(encoded string) string {
	freqmap := make(map[rune]int)
	for _, letter := range encoded {
		freqmap[letter]++
	}

	freqslice := make([]LetterFrequency, 0, len(freqmap))
	for letter, frequency := range freqmap {
		freqslice = append(freqslice, LetterFrequency{letter, frequency})
	}

	sort.Slice(freqslice, func(i, j int) bool {
		return freqslice[i].frequency > freqslice[j].frequency
	})

	// здесь решил не сразу записывать в ответ, чтобы
	// не было много перезаписей слайса
	end := 0
	for e, freq := range freqslice {
		if freq.letter == '_' {
			end = e
			break
		}
	}

	ans := make([]rune, 0, end)
	for i := 0; i < end; i++ {
		ans = append(ans, freqslice[i].letter)
	}

	return string(ans)
}
