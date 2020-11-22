### 排序算法

#### 冒泡排序

    func bubbleSort(nums []int) {
        length := len(nums)
        for i := 0; i < length-1; i++ {
            for j := i; j < length-1; j++ {
                if nums[j] > nums[j+1] {
                    nums[j], nums[j+1] = nums[j+1], nums[j]
                }
            }
        }
        return nums
    }

#### 插入排序
    func insertionSort(nums []int)[]int {
        length := len(nums)
        preIndex := 0
        current := 0
        for i := 1; i < length; i++ {
            preIndex = i - 1
            current = nums[i]
            for preIndex >= 0 && current < nums[preIndex] {
                nums[preIndex+1] = nums[preIndex]
                preIndex--
            }
            nums[preIndex+1] = current
        }
        return nums
    }


#### 选择排序

    func selectionSort(nums []int) {
        length := len(nums)
        for i := 0; i < length-1; i++ {
            minIndex := i
            for j := i + 1; j < length; j++ {
                if nums[j] < nums[minIndex] {
                    minIndex = j
                }
            }
            nums[i], nums[minIndex] = nums[minIndex], nums[i]
        }
        return nums
    }

#### 归并排序 
    func mergeSort(nums []int) []int {
        n := len(nums)
        if n < 2 {
            return nums
        }
        key := n / 2
        left := mergeSort(nums[0:key])
        right := mergeSort(nums[key:])
        return merge(left, right)
    }
    func merge(left, right []int) []int {
        newArr := make([]int, len(left)+len(right))
        i, j, index := 0, 0, 0
        for {
            if left[i] > right[j] {
                newArr[index] = right[j]
                index++
                j++
                if j == len(right) {
                    copy(newArr[index:], left[i:])
                    break
                }
            } else {
                newArr[index] = left[i]
                index++
                i++
                if i == len(left) {
                    copy(newArr[index:], right[j:])
                    break
                }
            }
        }
        return newArr
    }