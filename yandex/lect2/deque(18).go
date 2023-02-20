package lect2

import "fmt"

type Deque struct {
	deque  []int
	length int
}

func (s *Deque) Push_back(el int) {
	s.deque = append(s.deque, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
}

func (s *Deque) Pop_back() {
	if s.length == 0 {
		fmt.Println("error")
		return
	}
	s.deque = s.deque[:s.length-1]
	s.length--
}

func (s *Deque) Back() {
	if s.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(s.deque[s.length-1])
}

func (q *Deque) Push_front(el int) {
	q.deque = append([]int{el}, q.deque...)
	q.length++
	fmt.Println("ok")
}

func (q *Deque) Pop_front() {
	if q.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(q.deque[0])
	q.deque = q.deque[1:]
	q.length--
}

func (q *Deque) Front() {
	if q.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(q.deque[0])
}

func (q *Deque) Size() int {
	return q.length
}

func (q *Deque) Clear() {
	q.deque = []int{}
	q.length = 0
	fmt.Println("ok")
}

func Deque1() {
	MyStack := Deque{
		deque:  make([]int, 0),
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

		if str == "back" {
			MyStack.Back()
			continue
		}

		if str == "front" {
			MyStack.Front()
			continue
		}

		if str == "size" {
			fmt.Println(MyStack.Size())
			continue
		}

		if str == "pop_back" {
			MyStack.Pop_back()
			continue
		}

		if str == "pop_front" {
			MyStack.Pop_front()
			continue
		}

		if str == "clear" {
			MyStack.Clear()
			continue
		}

		if str == "push_back" {
			MyStack.Push_back(n)
		}
		if str == "push_front" {
			MyStack.Push_front(n)
		}

	}
}
