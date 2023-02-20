package lect1

import "fmt"

type BracketStack struct {
	stack  []string
	length int
}

func (s *BracketStack) Push(el string) string {
	s.stack = append(s.stack, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
	return "ok"
}

func (s *BracketStack) Pop() string {
	if s.length == 0 {
		return "error"
	}
	res := s.Back()
	s.stack = s.stack[:s.length-1]
	s.length--
	return res
}

func (s *BracketStack) Back() string {
	if s.length == 0 {
		return ""
	}
	return s.stack[s.length-1]
}

func (s *BracketStack) Size() int {
	return s.length
}

func Bracket() {
	var str string
	fmt.Scan(&str)
	myStack := BracketStack{
		stack: make([]string, 0),
	}
	if str == "" {
		fmt.Println("yes")
		return
	}
	for _, rune := range str {
		s := string(rune)
		if s == "(" || s == "{" || s == "[" {
			myStack.Push(s)
		} else if s == ")" && myStack.Size() != 0 {
			if myStack.Back() != "(" {
				fmt.Println("no")
				return
			}
			myStack.Pop()
		} else if s == "}" && myStack.Size() != 0 {
			if myStack.Back() != "{" {
				fmt.Println("no")
				return
			}
			myStack.Pop()
		} else if s == "]" && myStack.Size() != 0 {
			if myStack.Back() != "[" {
				fmt.Println("no")
				return
			}
			myStack.Pop()
		} else {
			fmt.Println("no")
			return
		}
	}
	// fmt.Println(myStack)
	if myStack.Size() != 0 {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}
