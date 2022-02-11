package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	冒泡排序 Bubble Sort O(n2)
	- 数组每两个相邻的数，如果前面比后面大，则两数交换。
	- 每躺排序完成后，则无需区减少一个数，有序区增加一个数。
	- 代码逻辑:
		- 总趟数: 数组长度 - 1  ---> len - 1
		- 当前趟数: for循环依次迭代 ---> i
		- 每趟比较数的长度: n - i - 1
			- 从第一个数开始，与后面的数依次比较，若大于后面的数，则两数交换，否则不变
*/

func BubbleSort(arrays []int)([]int) {
	// 循环所有趟数
	for i:=0; i < len(arrays) - 1; i ++ {
		// 优化，用于判断中间趟数是否有不需要交换的趟，有的话剩余的趟数则无需继续判断
		flag := false
		// 循环每一趟需要比较的个数
		for j := 0; j < len(arrays) - i - 1; j ++ {
			// 从第一个数开始，与后面的数依次比较，若大于后面的数，则两数交换，否则不变
			if arrays[j] > arrays[j + 1] {
				// 交叉赋值
				arrays[j], arrays[j + 1] = arrays[j + 1], arrays[j]
				flag = true
			}
		}
		if flag {
			return arrays
		}
	}
	return arrays
}

/*
	选择排序 Select Sort
	- 原理:
		- 每次从数组中获取最小的数，放到一个当前数组的第一个位置，以此类推
			- 第一个最小，放在第一个位置
			- 第二个最小，放在第二个位置
*/
func SelectSort(arrays []int)([]int){
	for i:=0; i < len(arrays) - 1; i ++ {
		// 初始化 最小值 索引 再次初始化的时候往下一个索引假设最小值
		minIndex := i
		// 从自身的下一个开始比较
		for j:=i + 1; j < len(arrays); j ++ {
			if arrays[j] < arrays[minIndex] {
				// 下一个值比当前的最小值小，则将其索引设置为本躺的最小值索引，以此类推
				minIndex = j
			}
		}
		// 循环结束后，则得到最小值的索引
		// 判断是否是当前值索引，若不是则将最小值与第一个值替换，否则跳过替换
		if minIndex != i {
			arrays[i], arrays[minIndex] = arrays[minIndex], arrays[i]
		}
	}
	return arrays
}

/*
	插入排序
*/

/*
	快速排序
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	i := 1
	arrays := make([]int, 0, 20)
	for i < 11 {
		// 获取1-20的随机数
		randNum := rand.Intn(20 - 1) + 1
		arrays = append(arrays, randNum)
		i ++
	}

	result := BubbleSort(arrays)
	fmt.Println("最后的结果是1:", result)

	result2 := SelectSort(arrays)
	fmt.Println("最后的结果是2:", result2)

}