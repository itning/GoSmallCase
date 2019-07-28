package main

import (
	"fmt"
)

func main() {
	data := []int{6, 7, 3, 5, 1, 2, 4}
	BubbleSort(data)
	fmt.Println(data)
}

// 冒泡排序 (稳定)
// 时间复杂度 O(n2)
func BubbleSort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		isChanged := false
		for j := 0; j < n-1-i; j++ {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}
