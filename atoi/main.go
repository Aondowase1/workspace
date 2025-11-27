package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	digits := []byte{}

	for n > 0 {
		d := n % 10
		digits = append(digits, byte('0'+d))
		n /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return sign + string(digits)
}
