package 

func PrintFish(s string, n int) {
	if n < 0 {
		Print("error")
	}
    if n%2 == 0 && n%3 ==  0 {
		Print("fish and chips")
	}
    if n%2 == 0 {
		Print("fish")
	}
    if n%3 == 0 {
		Print("chips")
	}
}