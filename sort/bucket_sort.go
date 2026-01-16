package main

func BucketSort(list []int, k int) {
	if len(list) == 0 {
		return
	}
	// 考虑到负数以及区间离0远的情况
	minVal, maxVal := list[0], list[0]
	for _, v := range list {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	// 防止所有数都是一个值
	if maxVal-minVal == 0 {
		return
	}
	// 初始化桶
	buckets := make([][]int, k)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	// 入桶
	for _, v := range list {
		// 把 [minVal, maxVal] 的数值区间，等比例映射到 [0, bucketCount) 的桶区间
		idx := (v - minVal) * (k - 1) / (maxVal - minVal)
		buckets[idx] = append(buckets[idx], v)
	}
	// 对每个桶内元素进行排序，这里使用插入排序，因为稳定
	for _, bucket := range buckets {
		InsertionSort(bucket)
	}
	// 合并区间
	i := 0
	for _, bucket := range buckets {
		for _, v := range bucket {
			list[i] = v
			i++
		}
	}
}
