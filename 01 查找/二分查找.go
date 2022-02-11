package main

import "fmt"

/*
	二分查找
		- 接收参数:
			- 被查找的数组
			- 查找的元素
*/
func BinarySearch(arrays[]int, searchNum int)(result int){
	// 数组左边的初始index
	left := 0
	// 获取数组的长度当做最右边的初始index
	right := len(arrays)
	// 中位数的初始索引
	mid := 0
	// 若左边的索引小于右边的索引，则开始循环
	for left < right {
		// 开始取半, 目的是为了继续从剩下的数取出中间数
		mid = int((left + right) / 2)
		if arrays[mid] < searchNum {
			left = mid + 1
		}else if arrays[mid] > searchNum {
			right = mid - 1
		}else {
			return mid  // 返回目标元素的index
		}
	}
	return 0
}

func main() {

	arraysNumber := make([]int, 0, 100000)
	for i:=1; i < 1000001; i ++ {
		arraysNumber = append(arraysNumber, i)
	}

	searchNumber := 9527

	midIndex := BinarySearch(arraysNumber, searchNumber)
	fmt.Println(midIndex)

}
