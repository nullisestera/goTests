package main

import (
	"fmt"
	"strconv"
)

func main() {
	evenAndOdds := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range evenAndOdds {
		if num%2 == 0 {
			fmt.Println(strconv.Itoa(num) + " is even")
		} else {
			fmt.Println(strconv.Itoa(num) + " is odd")
		}
	}
}
