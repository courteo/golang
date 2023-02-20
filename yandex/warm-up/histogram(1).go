package warm_up

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Histogram() {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			str += text
		} else {
			break
		}

	}
	symbols := map[string]int{}
	keys := []string{}
	maxCount := -1
	for _, i := range str {
		if string(i) == " " {
			continue
		}

		if _, ok := symbols[string(i)]; ok {
			symbols[string(i)]++
		} else {
			symbols[string(i)] = 1
			keys = append(keys, string(i))
		}

		if maxCount < symbols[string(i)] {
			maxCount = symbols[string(i)]
		}
	}

	sort.Strings(keys)
	for maxCount > 0 {
		for _, key := range keys {
			if symbols[key] >= maxCount {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		maxCount--
		fmt.Println()
	}

	for _, key := range keys {
		fmt.Print(key)
	}
}
