package leet

import "fmt"

//给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。其中第i个物品的重量为wt[i]，价值为val[i]，
//现在让你用这个背包装物品，最多能装的价值是多少？
//
//举个简单的例子，输入如下：
//N = 3, W = 4
//wt = [2, 1, 3]
//val = [4, 2, 3]
//算法返回 6，选择前两件物品装进背包，总重量 3 小于W，可以获得最大价值 6。

//最基础的01背包问题

//设dp数组：dp[i][w] 为对于当前第i个物品，容量为w ，可以装的最大容量是dp[i][w]
//状态转移方程：dp[i][w] = max(dp[i-1][w],dp[i-1][w - wt[i-1]] + val[i-1])

func knapsack(w, n int, wt, val []int) int {

	//go二维切片初始化不友好啊，还是说我使用方式不对呢
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}

	//因为状态转移方程中包含i-1,索性就从1开始
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 {
				//当前容量装不下第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			}
		}
	}
	return dp[n][w]
}

func TestKnapsack() {
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	result := knapsack(4, 3, wt, val)
	fmt.Printf("结果为%d", result)
}
