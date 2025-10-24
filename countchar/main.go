package main

func CountChar(s string, c rune) int {
	if s == "" {
		return 0
	}
	count := 0
	for _, char := range s {
		if char == c {
			count++
		}
	}
	return count
}
