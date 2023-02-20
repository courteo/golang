package lect1

import (
	"fmt"
)

type Stack struct {
	stack  []int
	length int
}

func (s *Stack) Push(el int) string {
	s.stack = append(s.stack, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
	return "ok"
}

func (s *Stack) Pop() {
	if s.length == 0 {
		fmt.Println("error")
		return
	}
	s.Back()
	s.stack = s.stack[:s.length-1]
	s.length--
}

func (s *Stack) Back() {
	if s.length == 0 {
		fmt.Println("error")
		return
	}
	fmt.Println(s.stack[s.length-1])
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Clear() {
	s.stack = []int{}
	s.length = 0
	fmt.Println("ok")
}

func Stack1() {
	MyStack := Stack{
		stack:  make([]int, 0),
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

		if str == "size" {
			fmt.Println(MyStack.Size())
			continue
		}

		if str == "pop" {
			MyStack.Pop()
			continue
		}

		if str == "clear" {
			MyStack.Clear()
			continue
		}

		if str == "push" {
			fmt.Println(MyStack.Push(n))
		}

	}
}
