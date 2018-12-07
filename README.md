# 使用go语言实现常见的算法


## [二分查找](https://github.com/g5niusx/algorithm-demo/blob/master/search/search.go)

对于有序的数组或者切片，通过与中间的数值进行大小比较，来实现快速查找指定的元素

## [选择排序](https://github.com/g5niusx/algorithm-demo/blob/master/search/search.go)

如果要对一个无序的数组或者切片进行排序，经常使用的方法是找出数组中最大的一个元素，然后放入到一个新的数组中，再次遍历，再找出一个最大的元素放入到数组中
一直重复，这种方法的时间复杂度为O(n*n)