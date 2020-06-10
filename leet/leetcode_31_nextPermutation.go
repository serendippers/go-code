package leet

import "sort"

func nextPermutation(nums []int)  {

	if len(nums) < 2 {
		return
	}

	var left int
	var right int
	for i := len(nums) -1; i>0; i -- {
		if nums[i] > nums[i -1] {
			left = i-1
			right = i
			for j := i; j<len(nums);j ++ {
				if nums[left] >= nums[j] {
					right = j -1
					break
				}
				if j == len(nums) -1 {
					right = j
				}
			}
			break
		}
		if i == 1 {
			sort.Ints(nums)
			return
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	sort.Ints(nums[left+1:])
}

