package go_container

import (
	"container/list"
	"fmt"
)

var linkedList *list.List

func createList()  {
	linkedList = list.New()
}

func playList()  {
	e1 := linkedList.PushBack("节点1")
	linkedList.PushBack("节点2")
	e2 := linkedList.PushFront("节点3")
	linkedList.PushBackList(linkedList)
	linkedList.PushFrontList(linkedList)
	linkedList.InsertBefore("节点4", e2)
	linkedList.InsertAfter("节点5", e2)
	linkedList.Remove(nil)
	linkedList.MoveAfter(e1, e2)
	linkedList.MoveBefore(e1, e2)
	linkedList.MoveToBack(e1)
	linkedList.MoveToFront(e2)

	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestGoList()  {
	createList()
	playList()
}