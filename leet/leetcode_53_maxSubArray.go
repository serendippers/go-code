package leet

//贪心
func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}
	return res
}

//动态规划
// f(x) = max{f(x), f(x) + f(x-1) } x > 1
// x= 1时，f(x)
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]< nums[i-1] + nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
