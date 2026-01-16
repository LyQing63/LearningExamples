package main

import (
	"reflect"
	"testing"
)

// SortFunc 统一的排序函数签名，添加新算法时实现此签名即可
type SortFunc func([]int) []int

func wrapQuickSortHorare(list []int) []int {
	if len(list) > 0 {
		QuickSort_Horare(list, 0, len(list)-1)
	}
	return list
}

func wrapQuickSortLomuto(list []int) []int {
	if len(list) > 0 {
		QuickSort_Lomuto(list, 0, len(list)-1)
	}
	return list
}

func wrapMergeSort(list []int) []int {
	if len(list) > 0 {
		MergeSort(list, 0, len(list)-1)
	}
	return list
}

func wrapBucketSort(list []int) []int {
	if len(list) > 0 {
		BucketSort(list, 10)
	}
	return list
}

// sortAlgorithms 注册所有排序算法，添加新算法只需在此处添加一行
var sortAlgorithms = map[string]SortFunc{
	"BubbleSort":       BubbleSort,
	"SelectSort":       SelectSort,
	"InsertionSort":    InsertionSort,
	"ShellSort":        ShellSort,
	"QuickSort_Horare": wrapQuickSortHorare,
	"QuickSort_Lomuto": wrapQuickSortLomuto,
	"MergeSort":        wrapMergeSort,
	"HeapSort":         HeapSort,
	"CountingSort":     CountingSort,
	"BucketSort":       wrapBucketSort,
}

var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{"空数组", []int{}, []int{}},
	{"单元素", []int{1}, []int{1}},
	{"两个元素-已排序", []int{1, 2}, []int{1, 2}},
	{"两个元素-逆序", []int{2, 1}, []int{1, 2}},
	{"普通数组", []int{5, 2, 8, 1, 9, 3}, []int{1, 2, 3, 5, 8, 9}},
	{"已排序数组", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"逆序数组", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"包含重复元素", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
	{"全部相同元素", []int{7, 7, 7, 7, 7}, []int{7, 7, 7, 7, 7}},
	{"包含负数", []int{-3, 5, -1, 0, 2, -8}, []int{-8, -3, -1, 0, 2, 5}},
}

// TestAllSortAlgorithms 遍历测试所有注册的排序算法
func TestAllSortAlgorithms(t *testing.T) {
	for algoName, sortFn := range sortAlgorithms {
		t.Run(algoName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					input := make([]int, len(tc.input))
					copy(input, tc.input)

					result := sortFn(input)

					if !reflect.DeepEqual(result, tc.expected) {
						t.Errorf("got %v, want %v", result, tc.expected)
					}
				})
			}
		})
	}
}

// BenchmarkSortAlgorithms 基准测试所有排序算法
func BenchmarkSortAlgorithms(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		data := generateRandomSlice(size)

		for algoName, sortFn := range sortAlgorithms {
			b.Run(algoName+"-"+itoa(size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					input := make([]int, len(data))
					copy(input, data)
					sortFn(input)
				}
			})
		}
	}
}

// generateRandomSlice 使用线性同余生成器生成伪随机切片，确保基准测试可重复
func generateRandomSlice(n int) []int {
	result := make([]int, n)
	seed := 12345
	for i := 0; i < n; i++ {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		result[i] = seed % 10000
	}
	return result
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}
