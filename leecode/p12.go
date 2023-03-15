package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	for _, v := range nums {
		var newSet [][]int
		for _, s := range ans {
			newSet = append(newSet, append(s, v))
		}
		for _, v := range newSet {
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {
	var s int
	fmt.Scan(&s)
	var d []int = make([]int, 0)
	for i := 0; i < s; i++ {
		var t int
		fmt.Scan(&t)
		d = append(d, t)
	}
	fmt.Println()
}
