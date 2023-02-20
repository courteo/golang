package lect2

import "fmt"

type HeapMax struct {
	heap Deque
}

func (h *HeapMax) Insert(el int) {
	h.heap.Push_back(el)
	i := h.heap.Size() - 1
	for i > 0 {

		prev := (i - 1) / 2
		// fmt.Println(i, prev, h.heap.deque)
		if h.heap.deque[prev] >= h.heap.deque[i] {
			temp := h.heap.deque[i]
			h.heap.deque[i] = h.heap.deque[prev]
			h.heap.deque[prev] = temp
			i = prev
		} else {
			break
		}
	}
}

func (h *HeapMax) Extract() {
	// fmt.Println("deque ", h.heap.deque)
	max := h.heap.deque[0]
	fmt.Print(max)
	if h.heap.Size() != 1 {
		fmt.Print(" ")
	}
	h.heap.deque[0] = h.heap.deque[h.heap.Size()-1]
	h.heap.Pop_back()
	// fmt.Println("new deque ", h.heap.deque)
	i := 0
	for i < h.heap.Size() {
		fisrtChild := 2*i + 1
		secondChild := 2*i + 2
		if fisrtChild < h.heap.Size() && secondChild < h.heap.Size() {
			if h.heap.deque[i] <= h.heap.deque[fisrtChild] && h.heap.deque[i] <= h.heap.deque[secondChild] {
				return
			} else {
				if h.heap.deque[fisrtChild] <= h.heap.deque[secondChild] {
					temp := h.heap.deque[fisrtChild]
					h.heap.deque[fisrtChild] = h.heap.deque[i]
					h.heap.deque[i] = temp
					i = fisrtChild
				} else {
					temp := h.heap.deque[secondChild]
					h.heap.deque[secondChild] = h.heap.deque[i]
					h.heap.deque[i] = temp
					i = secondChild
				}
			}
		} else if secondChild < h.heap.Size() {
			if h.heap.deque[i] <= h.heap.deque[secondChild] {
				return
			}
			temp := h.heap.deque[secondChild]
			h.heap.deque[secondChild] = h.heap.deque[i]
			h.heap.deque[i] = temp
			i = secondChild
		} else if fisrtChild < h.heap.Size() {
			if h.heap.deque[i] <= h.heap.deque[fisrtChild] {
				return
			}
			temp := h.heap.deque[fisrtChild]
			h.heap.deque[fisrtChild] = h.heap.deque[i]
			h.heap.deque[i] = temp
			i = fisrtChild
		} else {
			return
		}
	}
}

func Heap() {
	heap := HeapMax{
		heap: Deque{
			deque: make([]int, 0),
		},
	}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var k, m int
		fmt.Scan(&k)
		// fmt.Println("aa ", k, m)
		if k == 1 {
			heap.Extract()
		} else {
			fmt.Scanln(&m)
			heap.Insert(m)
			// fmt.Println(heap.heap.deque)
		}
	}
	// heap.Insert(2)
	// heap.Insert(5)
	// heap.Insert(4)
	// heap.Insert(11)
	// heap.Insert(6)
	// heap.Insert(8)
	// heap.Insert(25)
	// heap.Insert(12)
	// heap.Insert(20)
	// heap.Insert(3)
	// fmt.Println(heap.heap.deque)
	// heap.Extract()
	// fmt.Println(heap.heap.deque)
}
