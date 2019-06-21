package main

import "fmt"

func main() {
	//for i := 1488; i <= 1514; i++ {
	for i := rune('\u0400'); i <= rune('\u042f'); i++ {
		fmt.Printf("%d %c\n", i, i)
	}
}
