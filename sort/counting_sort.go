package main

func CountingSort(list []int) []int {
	if len(list) == 0 {
		return list
	}
	maxVal, minVal := list[0], list[0]
	for _, v := range list {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}
	offset := -minVal
	count := make([]int, maxVal-minVal+1)
	for _, v := range list {
		count[v+offset]++
	}
	// 前缀和定位，及大于等于i的个数
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	result := make([]int, len(list))
	// 保证稳定输出从后往前
	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		// 减1是因为从0开始计数，比如说比大于等于第一个数的有n个，那么索引就应该是0～n-1
		position := count[v+offset] - 1
		result[position] = v
		count[v+offset]--
	}
	return result
}
