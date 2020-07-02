package leet

import (
	"fmt"
	"math"
)

//378. 有序矩阵中第K小的元素
//没想出来，看了官方题解，惭愧

func kthSmallest(matrix [][]int, k int) int {

	//矩阵高
	h := len(matrix)
	//矩阵宽
	len := len(matrix[0])

	if h == 1 {
		return matrix[0][k-1]
	}

	var ltNum int
	l := matrix[0][0]
	r := matrix[len-1][h-1]
	result := -math.MaxInt32

	for r >= l {
		result = 0
		ltNum = 0
		middle := (l + r) / 2
		i, j := h-1, 0
		for i >= 0 && j <= len-1 {
			if middle >= matrix[i][j] {
				result = max(matrix[i][j], result)
				if j <= len-1 {
					ltNum++
				}
				j++
			} else {
				if i >= 0 {
					ltNum += j
				}
				i--
			}
		}
		if i > 0 {
			ltNum += i * len
		}
		if ltNum == k {
			return result
		}

		if ltNum > k {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}

	return result
}

func Test379() {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	result := kthSmallest(matrix, 8)
	fmt.Println(result)
}
