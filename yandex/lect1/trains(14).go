package lect1

import (
	"bufio"
	"fmt"
	"os"
)

type TrainStack struct {
	stack  []int
	length int
}

func (s *TrainStack) Push(el int) string {
	s.stack = append(s.stack, el)
	// fmt.Println("stack ", s.stack, el)
	s.length++
	return "ok"
}

func (s *TrainStack) Pop() (int, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return 0, err
	}
	res, _ := s.Back()
	s.stack = s.stack[:s.length-1]
	s.length--
	return res, nil
}

func (s *TrainStack) CheckIsBack(el int) bool {
	last, err := s.Back()
	if err != nil {
		return false
	}
	if last == el {
		return true
	}
	return false
}

func (s *TrainStack) Back() (int, error) {
	if s.length == 0 {
		err := fmt.Errorf("error")
		return 0, err
	}
	return s.stack[s.length-1], nil
}

func (s *TrainStack) Size() int {
	return s.length
}

func Train() {
	IsExist := make(map[int]int)
	deadEnd := TrainStack{
		stack: make([]int, 0),
	}

	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		if k == 1 {
			IsExist[1] = 1
			temp := 2
			for deadEnd.CheckIsBack(temp) {
				IsExist[temp] = 1
				deadEnd.Pop()
				temp++
			}
			// fmt.Println(k, temp, deadEnd)
		} else if _, ok := IsExist[k-1]; !ok {
			deadEnd.Push(k)
		} else {
			IsExist[k] = 1
			temp := k + 1
			for deadEnd.CheckIsBack(temp) {
				IsExist[temp] = 1
				deadEnd.Pop()
				temp++
			}
			// fmt.Println(k, temp, deadEnd)
		}
	}
	// fmt.Println(IsExist)
	// fmt.Println(deadEnd)
	for i := 1; i < n+1; i++ {
		if _, ok := IsExist[i]; !ok {
			last, _ := deadEnd.Pop()
			if last != i {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
