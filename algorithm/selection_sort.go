package main

import (
	"fmt"
)

func main() {
	data := []int{6, 7, 3, 5, 1, 2, 4}
	SelectionSort(data)
	fmt.Println(data)
}

// 选择排序 (稳定)
// 时间复杂度 O(n2)
func SelectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		// 假定首元素为最小元素
		min := i
		for j := min + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		// 将此次筛选出的最小元素放入最左边
		if min != i {
			data[min], data[i] = data[i], data[min]
		}
	}
}
