package week_01

// 利用语言的特性
func rotate(nums []int, k int) {
	k %= len(nums)
	ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	nums = append(nums[:0], ans...)

}
func rotate2(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
