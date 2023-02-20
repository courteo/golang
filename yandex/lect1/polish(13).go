package lect1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PolishStack struct {
	stack  []int
	length int
}

func (s *PolishStack) Push(el int) string {
	s.stack = append(s.stack, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
	return "ok"
}

func (s *PolishStack) Pop() (int, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return 0, err
	}
	res, _ := s.Back()
	s.stack = s.stack[:s.length-1]
	s.length--
	return res, nil
}

func (s *PolishStack) Back() (int, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return 0, err
	}
	return s.stack[s.length-1], nil
}

func (s *PolishStack) Size() int {
	return s.length
}

func Polish() {
	myStack := PolishStack{
		stack: make([]int, 0),
	}
	consoleReader := bufio.NewReader(os.Stdin)

	input, _ := consoleReader.ReadString('\n')
	var str string
	for _, rune := range input {
		str = string(rune)
		if str != " " {
			if str != "*" && str != "-" && str != "+" {
				n, err := strconv.Atoi(str)
				if err == nil {
					myStack.Push(n)
				}

			} else {
				second, _ := myStack.Pop()
				first, _ := myStack.Pop()
				if str == "*" {
					myStack.Push(first * second)
				} else if str == "-" {
					myStack.Push(first - second)
				} else {
					myStack.Push(first + second)
				}
				// fmt.Println(first, second, str, myStack)
			}
		}

	}
	res, _ := myStack.Pop()
	fmt.Println(res)
}
