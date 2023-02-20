package lect2

import (
	"fmt"
)

type QueueForDrunk struct {
	queue  []int
	length int
}

func (q *QueueForDrunk) Push(el int) {
	q.queue = append(q.queue, el)
	q.length++
}

func (q *QueueForDrunk) Pop() (int, error) {
	if q.length == 0 {
		return 0, fmt.Errorf("error")
	}
	last, _ := q.Front()
	q.queue = q.queue[1:]
	q.length--
	return last, nil
}

func (q *QueueForDrunk) Front() (int, error) {
	if q.length == 0 {
		return 0, fmt.Errorf("error")
	}
	return q.queue[0], nil
}

func (q *QueueForDrunk) Size() int {
	return q.length
}

func Drunk() {
	fisrt := QueueForDrunk{
		queue: make([]int, 0),
	}
	for i := 0; i < 5; i++ {
		var n int
		if i == 4 {
			fmt.Scanln(&n)
		} else {
			fmt.Scan(&n)
		}

		fisrt.Push(n)
	}
	// fmt.Println("qweqw ", fisrt)
	second := QueueForDrunk{
		queue: make([]int, 0),
	}
	for i := 0; i < 5; i++ {
		var n int
		if i == 4 {
			fmt.Scanln(&n)
		} else {
			fmt.Scan(&n)
		}
		second.Push(n)
	}
	// fmt.Println("asdas ", second)
	res := 0
	for fisrt.Size() != 0 && second.Size() != 0 && res != 1000000 {
		res++
		cardFirst, _ := fisrt.Pop()
		cardSecond, _ := second.Pop()
		if cardFirst == 0 && cardSecond == 9 {
			fisrt.Push(cardFirst)
			fisrt.Push(cardSecond)
		} else if cardFirst == 9 && cardSecond == 0 {
			second.Push(cardFirst)
			second.Push(cardSecond)
		} else if cardFirst > cardSecond {
			fisrt.Push(cardFirst)
			fisrt.Push(cardSecond)

		} else {
			second.Push(cardFirst)
			second.Push(cardSecond)
		}
		// fmt.Println(fisrt, second, res)
	}
	if fisrt.Size() != 0 {
		fmt.Println("first ", res)
	} else if second.Size() != 0 {
		fmt.Println("second ", res)
	} else {
		fmt.Println("botva")
	}
}
