package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		if n == 2 {
			return true
		}
		if n%2 == 0 {
			return false
		}
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				return false
			}
			return true
		}

		return true
	}

	for n := nb; n >= 2; n-- {
		if isPrime(n) {
			return n
		}
	}

	return 0
}
