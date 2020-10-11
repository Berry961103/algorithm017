package week_01

//   方法一 利用语言的特性 但利用了额外的内存
func rotate(nums []int, k int) {
	nLen := len(nums)
	if nLen <= 1 {
		return
	}

	k %= nLen
	ans := append(nums[nLen-k:], nums[:nLen-k]...)
	nums = append(nums[:0], ans...)

}

//  三次翻转
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
