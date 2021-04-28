package main

import "fmt"

func main() {
	parity([]int{
		2, 4, 0, 100, 4, 11, 2602, 36,
	})
	parity([]int{
		160, 3, 1719, 19, 11, 13, -21,
	})
	parity([]int{
		11, 13, 15, 19, 9, 13, -21,
	})
}

func parity(x []int) {
	odd := 0
	even := 0
	oddval := 0
	evenval := 0
	for _, n := range x {
		if n%2 == 0 {
			evenval = n
			even += 1
		} else {
			oddval = n
			odd += 1
		}
	}
	if odd < even && odd > 0 {
		fmt.Println(oddval)
		return
	}
	if even < odd && even > 0 {
		fmt.Println(evenval)
		return
	}
	fmt.Println("false")
}
