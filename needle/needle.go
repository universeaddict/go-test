package main

import "fmt"

func main() {
	needle([]string{
		"red", "blue", "yellow", "black", "grey",
	}, "blue")
}

func needle(x []string, y string) {
	for i, v := range x {
		if v == y {
			fmt.Println(i)
		}
	}
}
