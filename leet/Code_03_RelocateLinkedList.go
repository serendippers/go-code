package leet

import "fmt"

/**
示例：1-> 2 -> 3 -> 4 -> 5 -> 6 调整后 1 -> 4-> 2 -> 5 -> 3 -> 6
示例：1-> 2 -> 3 -> 4 -> 5 -> 6 -> 7 调整后 1 -> 4-> 2 -> 5 -> 3 -> 6 -> 7
 */
type Node struct {
	Val int
	Next *Node
}

/**
快慢指针,把链表分成两部分
*/
func RelocateLinkedList(head *Node) {
	if head == nil || head.Next == nil {
		return
	}
	mid := head
	right := head.Next

	show(head)
	//快慢指针，这个题边界有点难搞，痛苦😖
	for right.Next != nil && right.Next.Next != nil {
		mid = mid.Next
		right = right.Next.Next
	}
	right = mid.Next
	mid.Next = nil
	show(head)
	show(right)
	mergeLR(head, right)
}


/**
把两个链表合并
*/
func mergeLR(left, right *Node)  {
	var next *Node
	for left.Next != nil {
		next = right.Next
		right.Next = left.Next
		left.Next = right
		left = right.Next
		right = next
	}
	left.Next = right
}

func TestRelocateLinkedList()  {
	head := &Node{1, nil}
	temp := head
	for i:=2;i<8;i++  {
		head.Next = &Node{i,nil}
		head = head.Next
	}

	show(temp)

	RelocateLinkedList(temp)
	show(temp)
}

func show(head *Node)  {
	for head!=nil  {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}


