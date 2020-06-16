package leet

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(dp); i++ {
		// 取前面长度最长并且尾部数字小于nums[i]的
		dp[i] = 1 // 最小为1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
