package lect1

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	value int
	key   int
}

type PairStack struct {
	stack  []pair
	length int
}

func (s *PairStack) Push(el pair) string {
	s.stack = append(s.stack, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
	return "ok"
}

func (s *PairStack) Pop() (pair, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return pair{}, err
	}
	res, _ := s.Back()
	s.stack = s.stack[:s.length-1]
	s.length--
	return res, nil
}

func (s *PairStack) Back() (pair, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return pair{}, err
	}
	return s.stack[s.length-1], nil
}

func (s *PairStack) Size() int {
	return s.length
}

func TheNearestSmaller() {
	myStack := PairStack{
		stack: make([]pair, 0),
	}

	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		fmt.Fscan(in, &number)
		if myStack.Size() == 0 {
			myStack.Push(pair{
				key:   i,
				value: number,
			})
			continue
		}

		last, _ := myStack.Back()

		for last.value > number {
			myStack.Pop()
			res[last.key] = i
			last, _ = myStack.Back()
		}
		myStack.Push(pair{
			key:   i,
			value: number,
		})
		// fmt.Println("stack ", myStack)
	}

	for myStack.Size() != 0 {
		last, _ := myStack.Pop()
		res[last.key] = -1
	}

	for i := 0; i < n; i++ {
		fmt.Print(res[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}
