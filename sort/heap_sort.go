package main

import (
	"container/heap"
)

type SimpleHeap []int

func (h SimpleHeap) Len() int {
	return len(h)
}

func (h SimpleHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h SimpleHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SimpleHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *SimpleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSort(list []int) []int {
	h := &SimpleHeap{}
	heap.Init(h)
	for i := range list {
		heap.Push(h, list[i])
	}
	for i := 0; i < len(list); i++ {
		list[i] = heap.Pop(h).(int)
	}
	return list
}

// 自己实现一个堆
// parent: i, left: 2*i+1, right: 2*i+2
type MyHeap struct {
	data []int
}

func (h MyHeap) Len() int {
	return len(h.data)
}

func (h *MyHeap) up(i int) {
	if i <= 0 {
		return
	}
	parent := (i - 1) / 2
	if h.data[i] < h.data[parent] {
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		h.up(parent)
	}
}

func (h *MyHeap) down(i int) {
	left := 2*i + 1
	right := 2*i + 2
	lowest := i
	if left < h.Len() && h.data[left] < h.data[lowest] {
		lowest = left
	}
	if right < h.Len() && h.data[right] < h.data[lowest] {
		lowest = right
	}
	if lowest == i {
		return
	}
	h.data[i], h.data[lowest] = h.data[lowest], h.data[i]
	h.down(lowest)
}

func (h *MyHeap) Push(x int) {
	h.data = append(h.data, x)
	h.up(h.Len() - 1)
}

func (h *MyHeap) Pop() int {
	if len(h.data) == 0 {
		return 0
	}
	top := h.data[0]
	last := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	if h.Len() > 0 {
		h.data[0] = last
		h.down(0)
	}
	return top
}
