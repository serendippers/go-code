package leet


func Massage(nums []int) int {
	s := len(nums)
	if s == 0 {
		return 0
	}

	if s == 1 {
		return nums[0]
	}

	dp := make([]int, s)

	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < s; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[s-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}



//2



func OddEvenList(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	s1 := head
	s2 := head.Next.Next
	for s1 != nil && s2 != nil {
		index := s1.Next
		s1.Next = s2
		index.Next = s2.Next
		s := s2
		if s2.Next != nil{
			s = s2.Next.Next
		} else {
			s = nil
		}
		s2.Next = index
		s1 = s1.Next
		s2 = s
	}
	return head
}



