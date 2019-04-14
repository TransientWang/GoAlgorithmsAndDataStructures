package DynamicPlanning

import "math"

/*
	name:wfy
	516. 最长回文子序列
*/
func LongestPalindromeSubseq(s string) int {
	size := len(s)

	if size == 0 || size == 1 {
		return size
	}
	var dp = make([][]int, size) //不允许通过运行时计算得到的值来生成数组，只能只用切片
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}
	for i := size - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[0][size-1]

}
