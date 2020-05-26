package leet

/**
规定1是丑数，其他的数如果只含有2或3或5的因子，那么这个数也是丑数。
比如依次的丑数为:1,2,3,5,6,8,9,10,12,15...
求第n个丑数
*/
func UglyNumber1(index int) int {

	help := make([]int, index)

	i2 := 0
	i3 := 0
	i5 := 0
	help[0] = 1
	index = 1

	for index < len(help) {

		help[index] = intMin(help[i2]*2, intMin(help[i3]*3, help[i5]*5))

		if help[index] == help[i2]*2 {
			i2++
		}
		if help[index] == help[i3]*3 {
			i3++
		}

		if help[index] == help[i5]*5 {
			i5++
		}
		index++
	}
	return help[index-1]

}

/**
自己实现一个取最小值
*/
func intMin(i1, i2 int) int {

	if i2 < i1 {
		i1 = i2
	}
	return i1
}
