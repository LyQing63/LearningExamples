package main

// 空间复杂度 o(n)
func BubbleSort(list []int) []int {
	n := len(list)
	for i := 0; i < n; i++ {
		flag := false
		for j := n - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return list
}
