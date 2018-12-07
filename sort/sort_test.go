package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	ints := make([]int, 3)
	ints[0] = 3
	ints[1] = 1
	ints[2] = 2
	t.Logf("排序后结果:%v\n", SelectionSort(ints))
}
