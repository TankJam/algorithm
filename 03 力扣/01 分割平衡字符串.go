package main

import "fmt"

func main() {
	str := "RRLLRLRRRLLLRRRRLLLLRLRRLLRRLLRRRLLLL"

	num := 0
	result := 0
	for _, chr := range str {
		if chr == 'R' {
			num++
		} else {
			num--
		}
		if num == 0 {
			result ++
		}
	}

	fmt.Println(fmt.Sprintf("result: %d", result))
}


/*
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
*/