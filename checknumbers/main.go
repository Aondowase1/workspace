package main

import "github.com/01-edu/z01"

func CheckNumbers(arg string) bool {
	for _, r := range arg {
		if r := '0' && r <= '9' {
			return true
		}
	}
     return false

}


