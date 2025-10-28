package main

func HashCode(dec string) string {
	size := len(dec)
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		hashchar := (int(dec[i]) + size) % 127
		if hashchar < 33 {
			hashchar += 33
		}
		result[i] = byte(hashchar)
	}
	return string(result)
}
