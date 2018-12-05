package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := make([]int, 20)
	goal := 100
	for i := 0; i < 20; i++ {
		data[i] = i;
	}
	index := BinarySearch(data, goal)
	t.Logf("%v在%v的位置是%v\n", goal, data, index)
}
