package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	narcissistic(153)
	narcissistic(111)
}

func narcissistic(i int) bool {
	n := strconv.Itoa(i)
	len := len(n)
	var count float64
	for _, x := range n {
		xstr := string(x)
		xint, _ := strconv.Atoi(xstr)
		sum := math.Pow(float64(xint), float64(len))
		count += sum
	}
	if count != float64(i) {
		fmt.Println("false")
		return false
	}
	fmt.Println("True")
	return true
}
