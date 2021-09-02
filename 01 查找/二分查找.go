package main

import "fmt"

// BinarySearch - 二分查找 （时间复杂度: O(logn) ）
func BinarySearch(arrays []int, searchNumber int) int {
	left := 0
	right := len(arrays)
	mid := 0
	for left < right {
		// 1.每次循环都取数组中间索引
		mid = int((left + right) / 2) // 地板除
		if arrays[mid] == searchNumber {
			return mid
		} else if arrays[mid] > searchNumber {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

func main() {
	capNum := 1000000
	s1 := make([]int, 0, capNum)
	for i := 1; i <= capNum; i++ {
		s1 = append(s1, i)
	}
	result := BinarySearch(s1, 9527)
	fmt.Println(s1[result])
}
