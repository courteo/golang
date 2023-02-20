package warm_up

import (
	"bufio"
	"fmt"
	"os"
)

func GoodStr() {
	letters := map[string]int{
		"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0,
		"n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0,
	}
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var letter rune
	letter = 'a'
	for i := 0; i < n; i++ {
		var l int
		fmt.Fscan(in, &l)
		letters[string(letter)] = l
		letter++
	}
	count := 0
	//prev:= ""

	for {
		k := 0
		IsExistedPrevLetter := false
		letter = 'a'
		for i := 0; i < n; i++ {
			str := string(letter)
			if letters[str] != 0 {
				//letters[str]--
				temp := letters[string(letter-1)]
				if IsExistedPrevLetter == true {
					letters[string(letter-1)] = 0
					if letters[string(letter+1)] == 0 {
						letters[str] -= 0
						count += temp
					}
					//letters[str] -= temp
					count += temp
				} else {
					if str != "a" {
						letters[str] = 0

					}
					IsExistedPrevLetter = true
					k++
				}
			} else {
				k++
				IsExistedPrevLetter = false
			}
			fmt.Println(count, IsExistedPrevLetter, string(letter), letters)
			letter++
		}
		if k == n {
			break
		}
	}
	fmt.Println(count)
	//fmt.Println(count, letters)
}
