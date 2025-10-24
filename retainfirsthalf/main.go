package main

import (
	"fmt"
	"string"
)

func PrintFirstHalf(str string) {
	newstring := string.Trimspace(str)
	lenght := len(newstring)
	firsthalf := lenght / 2

	for i, val := range newstring {
		fmt.Print(string(val))
		if i == firsthalf-1 {
			break
		}
	}
}
