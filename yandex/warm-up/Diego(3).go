package warm_up

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Diego() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	stickers := map[int]int{}
	keys := []int{}
	length := 0
	for i := 0; i < n; i++ {
		var sticker int
		fmt.Fscan(in, &sticker)
		if _, ok := stickers[sticker]; !ok {
			stickers[sticker]++
			if length == 0 {
				keys = append(keys, sticker)
			} else {
				keys = insertSorted(keys, sticker)
			}

			length++

		}
	}
	//fmt.Println(keys)
	var k int
	fmt.Fscan(in, &k)
	for i := 0; i < k; i++ {
		var collector int
		fmt.Fscan(in, &collector)
		//fmt.Println(binarySearch(keys, collector, length, 0, length-1))
		index := binarySearch(keys, collector, 0, length-1)
		if keys[index] < collector {
			fmt.Println(length)
		} else {
			fmt.Println(index)
		}
	}
}

func insertAt(data []int, i int, v int) []int {
	if i == len(data) {
		return append(data, v)
	}
	data = append(data[:i+1], data[i:]...)
	data[i] = v
	return data
}

func insertSorted(data []int, v int) []int {
	i := sort.Search(len(data), func(i int) bool { return data[i] >= v })
	//fmt.Println("i ", i, v, data)
	return insertAt(data, i, v)
}

func binarySearch(keys []int, collector, begin, end int) int {
	for begin < end {
		middle := (end-begin)/2 + begin
		if keys[middle] >= collector {
			end = middle
		} else {
			begin = middle + 1
		}
	}
	return begin

	//fmt.Println(keys, collector, index, begin, end)
	//middle := (end-begin)/2 + begin
	//if end > begin {
	//	if collector <= keys[middle] {
	//		return binarySearch(keys, collector, middle, begin, middle)
	//	} else {
	//		return binarySearch(keys, collector, end+1, middle+1, end)
	//	}
	//} else if end == begin {
	//	fmt.Println("special: ", keys[end], collector, index)
	//	if collector <= keys[end] {
	//		return index - 1
	//	} else {
	//		return index
	//	}
	//} else {
	//	return index
	//}
}
