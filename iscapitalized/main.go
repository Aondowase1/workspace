package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	word := ""
	var words []string

	for _, r := range s {
		if r == ' ' || r == '\t' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		return false
	}

	for _, w := range words {
		first := w[0]

		if first >= 'a' && first <= 'z' {
			return false
		}

	}

	return true
}
