package leet

//回溯算法(深度优先的递归)
func generateParenthesis(n int) []string {

	var results []string
	if n < 1 {
		return results
	}
	//左括号剩余数量（因为第一个必须为左括号，所以根节点先使用一个）
	left := n-1
	//右括号剩余数量
	right := n


	str := "("
	doJoin(left, right, str, &results)


	return results
}

func doJoin(left, right int, str string, results *[]string) {

	if left> 0 {
		doJoin(left-1, right, str + "(", results)
	}

	if left == 0 && right> 0 {
		for right > 0 {
			str += ")"
			right--
		}
	}
	if right > 0 && right > left {
		doJoin(left, right-1, str + ")", results)
	}

	if left==0 && right==0 {
		*results = append(*results,  str)
	}
}

