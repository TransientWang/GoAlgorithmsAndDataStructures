package greedy

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
