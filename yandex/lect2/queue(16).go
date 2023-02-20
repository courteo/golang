package lect2

import "fmt"

type Queue struct {
	queue  []int
	length int
}

func (q *Queue) Push(el int) {
	q.queue = append(q.queue, el)
	q.length++
	fmt.Println("ok")
}

func (q *Queue) Pop() {
	if q.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(q.queue[0])
	q.queue = q.queue[1:]
	q.length--
}

func (q *Queue) Front() {
	if q.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(q.queue[0])
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) Clear() {
	q.queue = []int{}
	q.length = 0
	fmt.Println("ok")
}

func Queue1() {
	MyQueue := Queue{
		queue:  make([]int, 0),
		length: 0,
	}
	for {
		var str string
		var n int
		fmt.Scanf("%s %d", &str, &n)
		if str == "exit" {
			fmt.Println("bye")
			return
		}

		if str == "front" {
			MyQueue.Front()
			continue
		}

		if str == "size" {
			fmt.Println(MyQueue.Size())
			continue
		}

		if str == "pop" {
			MyQueue.Pop()
			continue
		}

		if str == "clear" {
			MyQueue.Clear()
			continue
		}

		if str == "push" {
			MyQueue.Push(n)
		}

	}
}
