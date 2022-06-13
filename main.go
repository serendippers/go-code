package main

import (
	"fmt"
	"go-code/leet"
)

func main() {
	//Mutex，
	//go_sync.GoMutex()
	//var sliceBase = go_base.SliceBase{}
	//sliceBase.NewSlice()
	//leet.TestKnapsack()
	//go_container.TestGoList()
	//nums := []int{2,1,4,5,3,1,1,3}
	//res := leet.Massage(nums)
	//fmt.Println(res)

	head :=leet.ListNode{2,nil}
	head.Next = &leet.ListNode{1, nil}
	head.Next.Next = &leet.ListNode{3, nil}
	head.Next.Next.Next = &leet.ListNode{5, nil}
	head.Next.Next.Next.Next = &leet.ListNode{6, nil}
	head.Next.Next.Next.Next.Next = &leet.ListNode{4, nil}
	//for s := head; &s != nil; {
	//	fmt.Println(s.Val)
	//	s= *s.Next
	//}
	leet.OddEvenList(&head)
	LinkPrint(&head)
}

func LinkPrint(cur *leet.ListNode) {
	for cur != nil {
		fmt.Printf("cur %v  , val %d \n", cur, cur.Val)
		cur = cur.Next
	}
}
