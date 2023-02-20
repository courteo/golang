package warm_up

import "fmt"

func MidTerm() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	var column, row int
	fmt.Scan(&row)
	fmt.Scan(&column)

	var temp, variant int
	if column == 2 {
		temp = 2 * row
	} else {
		temp = row*2 - 1
	}

	if temp <= k {
		variant = temp
	} else {
		variant = temp - (temp/k)*k
	}

	var count int
	if n-k*(n/k) >= variant {
		count = n/k + 1
	} else {
		count = n / k
	}

	//fmt.Println(temp, variant, count)
	if count == 1 {
		fmt.Println(-1)
		return
	}

	if temp-k > 0 {
		temp -= k
	} else {
		temp += k
		if temp > n {
			temp -= 2 * k
		}
	}
	// fmt.Println("temp ", temp)

	if temp%2 == 0 {
		fmt.Println(temp/2, 2)

	} else {
		fmt.Println((temp+1)/2, 1)

	}
}
