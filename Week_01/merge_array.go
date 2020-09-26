package week_01

func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right, aCount := m-1, n-1, m+n-1 //定义两个数组的最大的数的指针
	for left >= 0 && right >= 0 {          //先把num1与num2中其中一个全部填入
		if nums1[left] > nums2[right] {
			nums1[aCount] = nums1[left]
			left--
		} else {
			nums1[aCount] = nums2[right]
			right--
		}
		aCount--
	}
	//当num1未完全填充的时候
	for left >= 0 {
		nums1[aCount] = nums1[left]
		aCount--
		left--
	}
	//当num2未完全填充的时候
	for right >= 0 {
		nums1[aCount] = nums2[right]
		aCount--
		right--
	}
}
