package main

import "fmt"

func main() {
	data := []int{6, 7, 3, 5, 1, 2, 4}
	qSort(data)
	fmt.Println(data)
}

// 快速排序 (不稳定)
// 时间复杂度 O(n logn)
func qSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qSort(data[:head])
	qSort(data[head+1:])
}
