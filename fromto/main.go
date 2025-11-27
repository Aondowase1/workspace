package piscine

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	if from <= to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + fmt.Sprint(i)
			} else {
				result += fmt.Sprint(i)
			}

			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + fmt.Sprint(i)
			} else {
				result += fmt.Sprint(i)
			}

			if i != to {
				result += ", "
			}
		}
	}

	result += "\n"
	return result
}
