package main

import "testing"

var (
	nums = []int{2, 7, 11, 15}
)

func TestTwoSum(t *testing.T) {
	target := 22
	var result = TwoSum(nums, target)
	t.Logf("target: %d, return: %d + %d eq %d.\n", target, result[0], result[1], target)
}
