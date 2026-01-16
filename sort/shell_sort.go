package main

func ShellSort(list []int) []int {
	n := len(list)
	h := 1
	for h < n/3 {
		h = h*3 + 1
	}
	for h > 0 {
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				if list[j] < list[j-h] {
					list[j], list[j-h] = list[j-h], list[j]
				} else {
					break
				}
			}
		}
		h /= 3
	}
	return list
}
