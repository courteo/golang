package lect2

import "fmt"

func HeapSort() {
	heap := HeapMax{
		heap: Deque{
			deque: make([]int, 0),
		},
	}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		heap.Insert(k)
	}
	for i := 0; i < n; i++ {
		heap.Extract()
	}
}
