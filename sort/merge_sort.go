package main

func MergeSort(list []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2
	MergeSort(list, lo, mid)
	MergeSort(list, mid+1, hi)

	// 合并
	Merge(list, lo, mid, hi)
}

func Merge(list []int, lo int, mid int, hi int) {
	aux := make([]int, hi-lo+1)
	copy(aux, list[lo:hi+1])

	i, j := 0, mid-lo+1
	for k := lo; k <= hi; k++ {
		if i > mid-lo {
			list[k] = aux[j]
			j++
		} else if j > hi-lo {
			list[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			list[k] = aux[i]
			i++
		} else {
			list[k] = aux[j]
			j++
		}
	}
}
