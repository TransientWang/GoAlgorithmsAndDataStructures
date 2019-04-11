package mathquestion

import "math"

//191. 位1的个数
func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	res := uint32(1)

	for num > 2 {

		res += num % 2
		num = uint32(math.Floor(float64(num / 2)))
	}
	return int(res)
}
