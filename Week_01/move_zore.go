package week_01

func moveZero(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			//fmt.Println("11111111111", nums)
			j++
		}
	}
	return nums
}
