package main

func CountOpenBracket(list string) int {
	count := 0
	for _, v := range list {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		} else {
			panic("Wrong character")
		}
	}
	return count
}

func FindBasementPosition(list string) int {
	count := 0
	for idx, v := range list {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
			if count == -1 { return idx + 1 }
		} else {
			panic("Wrong character")
		}
	}
	return -1
}