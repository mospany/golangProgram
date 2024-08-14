package __twoSum

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
