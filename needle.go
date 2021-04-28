package main

import "fmt"

func main() {
	parity([]string{
		"red", "blue", "yellow", "black", "grey",
	}, "blue")
}

func parity(x []string, y string) {
	for i, v := range x {
		if v == y {
			fmt.Println(i)
		}
	}
}
