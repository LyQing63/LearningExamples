package main

// 时间复杂度 o(n^2)
// 空间复杂度 o(1)
func SelectSort(list []int) []int {
	n := len(list)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[minIndex] > list[j] {
				minIndex = j
				list[minIndex] = list[j]
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
	return list
}
