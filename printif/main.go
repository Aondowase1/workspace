package main

func PrintIf(s string) string {
	if s == "" {
		return "G\n"
	}
	if len(s) >= 3 {
		return "G\n"
	}
	return "invalid imput \n"
}
