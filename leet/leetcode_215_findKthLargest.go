package leet

import "math"

var heap []int

//两种思路，
//1. 给数组排序，取k位置。 
//2. 遍历数组，维护一个大小为k的空间存前k个数字，去最小（一开始想到的是维护一个链表方便更新元素，这里其实用一个小根堆是最合适的）
func findKthLargest(nums []int, k int) int {

	buildHeap(k)

	for _, v := range nums {
		push(v)
	}
	return heap[0]
}

func buildHeap(capacity int) {
	heap = make([]int, capacity)
	for i,_ := range heap {
		heap[i] = -math.MaxInt16
	}
}

func push(value int) {
	if value > heap[0] {
		heap[0] = value
	}
	down(0)
}

func down(i int) {
	//if i == len(heap)-1 {
	//	return
	//}
	left := i<<1 + 1
	right := i<<1 + 2
	var mid int
	if left > len(heap)-1 {
		return
	}

	if right <= len(heap)-1 {
		mid = min(left, right)
	} else {
		mid = left
	}


	if heap[i] < heap[mid]  {
		return
	} else {
		heap[mid], heap[i] = heap[i], heap[mid]
		down(mid)
	}
}

func min(a, b int) int {
	if heap[a] > heap[b] {
		return b
	}
	return a
}