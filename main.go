package main

import "go-code/leet"

func main() {

	//cache := leet.Constructor(2)
	//
	//cache.Put(1, 1);
	//cache.Put(2, 2);
	//cache.Get(1);       // 返回  1
	//cache.Put(3, 3);    // 该操作会使得关键字 2 作废
	//cache.Get(2);       // 返回 -1 (未找到)
	//cache.Put(4, 4);    // 该操作会使得关键字 1 作废
	//cache.Get(1);       // 返回 -1 (未找到)
	//cache.Get(3);       // 返回  3
	//cache.Get(4);       // 返回  4



	cache2 := leet.Constructor(1)
	cache2.Put(2,1)
	cache2.Get(2)
	cache2.Put(3,2)
}
