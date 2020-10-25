package week04

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5: // 给5块不用找
			five++
		case 10:
			if five >= 1 { // 给10块手上有大于等于一张5块的就找
				five--
				ten++
			} else { // 没有就结束
				return false
			}
		case 20:
			if five >= 1 && ten >= 1 { // 有一张10块且一张五块
				five--
				ten--
			} else if five >= 3 { // 有三张5块的
				five -= 3
			} else {
				return false
			}
		}

	}
	return true
}
