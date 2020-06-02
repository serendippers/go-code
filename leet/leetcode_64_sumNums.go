package leet

//求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
//不让乘除法——等差数列求和不能用；for while的迭代也不行，剩下的只有递归了
//ps：不喜欢这种题目
func sumNums(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}
