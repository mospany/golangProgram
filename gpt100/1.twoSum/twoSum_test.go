package __twoSum

import (
	"fmt"
	"testing"
)

type values struct {
	nums   []int
	target int
}

var cases = []values{
	{nums: []int{2, 7, 11, 15}, target: 9},
	{nums: []int{1, 3, 5, 7, 9}, target: 6},
	{nums: []int{1, 3, 5, 7, 9}, target: 20},
}

func TestTwoNum(t *testing.T) {
	for i, c := range cases {
		ret := twoSum(c.nums, c.target)
		fmt.Printf("Case_%d, nums: %v, target: %v, ret: %v.\n", i, c.nums, c.target, ret)
	}
}
