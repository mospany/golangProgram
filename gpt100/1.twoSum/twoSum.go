package __twoSum

/*
1. 两数之和
描述: 给定⼀个整数数组和⼀个⽬标值，找出数组中和为⽬标值的两个数的索引。
思路: 使⽤哈希表记录数字及其索引，遍历数组查找。
*/
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for idx, num := range nums {
		if i, ok := indexMap[target-num]; ok {
			return []int{idx, i}
		}
		indexMap[num] = idx
	}
	return nil
}
