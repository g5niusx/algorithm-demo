package sort

import "fmt"

// 选择排序
func SelectionSort(data []int) []int {
	// 创建新的切片来存放排序以后的切片
	for i := 0; i < len(data); i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			// 如果当前元素小于最小的元素，则赋值给最小的元素下标
			if data[j] < data[minIndex] {
				minIndex = j
				fmt.Println(data[minIndex])
			}
		}
		temp := data[i]
		data[i] = data[minIndex]
		data[minIndex] = temp
	}
	return data
}
