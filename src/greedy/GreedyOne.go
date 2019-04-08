package greedy

func Candy(ratings []int) int {
	/*
		135. 分发糖果
	*/
	var r []int = make([]int, len(ratings))
	r[0] = 1
	for i := 1; i < len(ratings); i++ {
		r[i] = 1
		if ratings[i] > ratings[i-1] && r[i] <= r[i-1] {
			r[i] = r[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && r[i] <= r[i+1] {
			r[i] = r[i+1] + 1
		}
	}
	sum := 0
	for _, val := range r {
		sum += val
	}
	return sum
}

func CanCompleteCircuit(gas []int, cost []int) int {
	/*
		134. 加油站
	*/
	tank, start, lst := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			lst += tank
			tank = 0
			start = i + 1
		}
	}
	if tank+lst < 0 {
		return -1
	}
	return start
}
