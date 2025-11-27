package piscine

func RepeatAlpha(s string) string {
	result := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' {
			count := int(ch-'a') + 1
			for j := 0; j < count; j++ {
				result = append(result, ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			count := int(ch-'A') + 1
			for j := 0; j < count; j++ {
				result = append(result, ch)
			}
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}
