package main

// 时间复杂度 o(n^2)
// 空间复杂度 o(1)
func InsertionSort(list []int) []int {
	n := len(list)
	// i是无序区的左边界
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			} else {
				break
			}
		}
	}
	return list
}
