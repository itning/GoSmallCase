package main

import "fmt"

func main() {
	data := []int{6, 7, 3, 5, 1, 2, 4}
	InsertSort(data)
	fmt.Println(data)
}

// 插入排序 (稳定)
// 时间复杂度 O(n2)
func InsertSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			} else {
				break
			}
		}
	}
}
