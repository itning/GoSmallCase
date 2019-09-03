package main

import "fmt"

// abc 能有多少种组合
func main() {
	input([]rune("abc"), 0)
}

func input(charArray []rune, index int) {
	if index == len(charArray) {
		fmt.Println(string(charArray))
	} else {
		for i := index; i < len(charArray); i++ {
			charArray[i], charArray[index] = charArray[index], charArray[i]
			input(charArray, index+1)
			charArray[i], charArray[index] = charArray[index], charArray[i]
		}
	}
}
