package week08

func reverseBits(num uint32) uint32 {

	var res uint32
	pointer := 31

	for num != 0 {
		res += (num & 1) << pointer
		pointer--
		num = num >> 1
	}
	return res
}
