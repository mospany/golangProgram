package main

import "fmt"

/*
  原理： 把值当hash key
*/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, b := range nums {
		m[b] = i
		if j, ok := m[target-b]; ok {
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	fmt.Println("hello")
}
