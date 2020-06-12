package go_base

import "fmt"

type SliceBase struct {

}

func (base *SliceBase) NewSlice()  {
	//nil 切片和nil相等，一般用来表示不存在的切片
	var a []int
	fmt.Printf("a []int type：%T, value: %v, len is %d, cap is %d\n", a, a, len(a), cap(a))
	if a == nil {
		fmt.Println("a 与nil相等")
	}

	//空切片，和nil不等，用来表示空集合
	var b = []int{}
	fmt.Printf("b = []int{} type：%T, value: %v, len is %d, cap is %d\n", b, b, len(b), cap(b))
	if b == nil {
		fmt.Println("b 与nil相等")
	}

	var c = []int{1, 2, 3}
	fmt.Printf("c = []int{1, 2, 3} type：%T, value: %v, len is %d, cap is %d\n", c, c, len(c), cap(c))

	//语法：切分1-2位置的元素，左闭右开
	var d = c[1:2]
	fmt.Printf("d = c[1:2] type：%T, value: %v, len is %d, cap is %d\n", d, d, len(d), cap(d))


	// 有2个元素的切片, len为2, cap为3
	e := c[0:2:cap(c)]
	fmt.Printf("e := c[0:2:cap(c)] type：%T, value: %v, len is %d, cap is %d\n", e, e, len(e), cap(e))

	// 有0个元素的切片, len为0, cap为3
	f := c[:0]
	fmt.Printf("f := c[:0] type：%T, value: %v, len is %d, cap is %d\n", f, f, len(f), cap(f))


	// 有3个元素的切片, len和cap都为3
	g := make([]int, 3)
	fmt.Printf("g := make([]int, 3) type：%T, value: %v, len is %d, cap is %d\n", g, g, len(g), cap(g))


	// 有2个元素的切片, len为2, cap为3
	h := make([]int, 2, 3)
	fmt.Printf("h := make([]int, 2, 3) type：%T, value: %v, len is %d, cap is %d\n", h, h, len(h), cap(h))


	// 有0个元素的切片, len为0, cap为3
	i := make([]int, 0, 3)
	fmt.Printf("i := make([]int, 0, 3) type：%T, value: %v, len is %d, cap is %d\n", i, i, len(i), cap(i))

	//和数组的最大不同是，切片的类型和长度信息无关，只要是相同类型元素构成的切片均对应相同的切片类型。

}
