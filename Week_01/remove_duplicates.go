package week_01

//  nums = [0,0,1,1,1,2,2,3,3,4],
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen < 2 {
		return numsLen
	}

	length := 1
	p := nums[0]
	for i := 1; i < numsLen; i++ {
		v := nums[i]
		if v != p {
			nums[length] = v
			p = v
			length++
		}
	}
	return length
}
