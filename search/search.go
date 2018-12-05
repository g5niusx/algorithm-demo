package search

// 二分查找
func BinarySearch(data []int, goal int) (index int) {
	return binarySearch(data, goal, 0, len(data)-1)
}

// 二分查找的逻辑实现,入参的切片必须是已经排序过的
func binarySearch(data []int, goal int, left int, right int) int {
	if left > right {
		return -1
	}
	// 求中间位置
	middle := (left + right) / 2
	if data[middle] == goal {
		return middle
	} else if data[middle] < goal {
		return binarySearch(data, goal, middle+1, right)
	} else {
		return binarySearch(data, goal, 0, middle-1)
	}
}
