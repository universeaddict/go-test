package main

import "fmt"

func main() {
	blueocean([]int{
		1, 2, 3, 4, 6, 10,
	}, []int{
		1,
	})
	blueocean([]int{
		1, 5, 5, 5, 5, 3,
	}, []int{
		5,
	})
}

func blueocean(x []int, y []int) {
	var result []int
	for _, v := range x {
		for _, n := range y {
			if v != n {
				result = append(result, v)
			}
		}
	}
	fmt.Println(result)
}
