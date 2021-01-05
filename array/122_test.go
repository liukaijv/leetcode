package array

import (
	"testing"
)

/**
买卖股票的最佳时机

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 

提示：

1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
贪心算法
只要今天价格小于明天价格就在今天买入然后明天卖出，时间复杂度O(n)
*/
func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	out := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			out += prices[i] - prices[i-1]
		}
	}

	return out
}

func TestMaxProfit(t *testing.T) {

	arr := []int{7, 1, 5, 3, 6, 4}
	expect := 7

	if maxProfit(arr) != expect {
		t.Fail()
	}

	arr = []int{1, 2, 3, 4, 5}
	expect = 4

	if maxProfit(arr) != expect {
		t.Fail()
	}

	arr = []int{7, 6, 4, 3, 1}
	expect = 0

	if maxProfit(arr) != expect {
		t.Fail()
	}

}

/**
动态规划
一般的动态规划题目思路三步走：

定义状态转移方程
给定转移方程初始值
写代码递推实现转移方程
*/
func maxProfitDP(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	//定义二维数组 dp[n][2]：
	//dp[i][0] 表示第 ii 天不持有可获得的最大利润；
	//dp[i][1] 表示第 ii 天持有可获得的最大利润（注意是第 ii 天持有，而不是第 ii 天买入）。
	dp := make([][2]int, len(prices))

	//对于第 00 天：
	//不持有： dp[0][0] = 0dp[0][0]=0
	//持有（即花了 price[0]price[0] 的钱买入）： dp[0][1] = -prices[0]dp[0][1]=−prices[0]
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	//定义状态转移方程：
	//不持有：dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
	//对于今天不持有，可以从两个状态转移过来：1. 昨天也不持有；2. 昨天持有，今天卖出。两者取较大值。
	//持有：dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
	//对于今天持有，可以从两个状态转移过来：1. 昨天也持有；2. 昨天不持有，今天买入。两者取较大值。
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxProfitDP(t *testing.T) {

	arr := []int{7, 1, 5, 3, 6, 4}
	expect := 7

	if maxProfitDP(arr) != expect {
		t.Fail()
	}

	arr = []int{1, 2, 3, 4, 5}
	expect = 4

	if maxProfitDP(arr) != expect {
		t.Fail()
	}

	arr = []int{7, 6, 4, 3, 1}
	expect = 0

	if maxProfitDP(arr) != expect {
		t.Fail()
	}

}
