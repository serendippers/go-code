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
	linkedList.PushBack("节点1")
	linkedList.PushBack("节点2")
	e := linkedList.PushFront("节点3")
	linkedList.PushBackList(linkedList)
	linkedList.PushFrontList(linkedList)
	linkedList.InsertBefore("节点4", e)
	linkedList.InsertAfter("节点5", e)

	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestGoList()  {
	createList()
	playList()
}