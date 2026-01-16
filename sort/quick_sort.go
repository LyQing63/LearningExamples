package main

// Horare 写法
func QuickSort_Horare(list []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition_Horare(list, lo, hi)
	QuickSort_Horare(list, lo, p)
	QuickSort_Horare(list, p+1, hi)
}

func partition_Horare(list []int, lo int, hi int) int {
	pivot := list[lo]
	i, j := lo-1, hi+1
	for {
		for {
			i++
			if list[i] >= pivot {
				break
			}
		}
		for {
			j--
			if list[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		list[i], list[j] = list[j], list[i]
	}
}

// Lomuto 写法
func QuickSort_Lomuto(list []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition_Lomuto(list, lo, hi)
	QuickSort_Lomuto(list, lo, p-1)
	QuickSort_Lomuto(list, p+1, hi)
}

func partition_Lomuto(list []int, lo int, hi int) int {
	pivot := list[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if list[j] < pivot {
			list[j], list[i] = list[i], list[j]
			i++
		}
	}
	list[i], list[hi] = list[hi], list[i]
	return i
}
