package week02

func towSum2(nums []int, target int) []int {
	hashTab := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		hashTab[target-nums[i]] = nums[i]
	}

	for j := 0; j < len(nums); j++ {
		if v, ok := hashTab[nums[j]]; ok {
			return []int{nums[j], v}
		}
	}
	return nil
}
