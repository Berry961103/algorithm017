package week_01

func towSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {

				return []int{i, j}
			}
		}
	}
	return nil
}

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
