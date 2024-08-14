package __twoSum

/*
给 定 一 个 整 数 数 组 nums 和 一 个 目 标 值 target，
请 你 在 该 数 组 中 找 出 和 为 目 标 值 的 那 两 个 整 数 ， 并 返 回 他 们 的 数 组 下 标 。
你 可 以 假 设 每 种 输 入 只 会 对 应 一 个 答 案 。 但 是 ， 你 不 能 重 复 利 用 这 个 数 组 中 同 样 的 元 素 。
示 例:给 定 nums = [2, 7, 11, 15], target = 9
因 为 nums[0] + nums[1] = 2 + 7 = 9
所 以 返 回 [0, 1]
*/
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
