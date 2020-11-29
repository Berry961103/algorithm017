package week09

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	d := make([]int, n+1)
	l := 1
	d[l] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[l] {
			l++
			d[l] = nums[i]
		} else {
			left := 1
			right := l
			pos := 0
			for left <= right {
				mid := (left + right) >> 1
				if nums[i] > d[mid] {
					pos = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return l
}
