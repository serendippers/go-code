package leet

//还应该再优化下代码，凌晨一点，先睡了，fuck you work！
func productExceptSelf(nums []int) []int {
	var arr = make([]int,len(nums))
	var result = make([]int,len(nums))
	for i, j := range nums{
		if i ==0 {
			arr[i] = j
			continue
		}
		arr[i] = j * arr[i-1]
	}


	for i := len(nums)-2; i>=0; i--{
		nums[i] = nums[i] * nums[i+1]
	}

	result[0]= nums[1]
	result[len(nums)-1]= arr[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		result[i]= arr[i-1] * nums[i+1]
	}
	return result
}
