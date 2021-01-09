package dp

import "testing"

/**
最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxSubArray(t *testing.T) {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6

	if maxSubArray(arr) != expect {
		t.Fail()
	}

}

func maxSubArray1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	result := nums[0]
	sum := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		result = max(result, sum)
	}

	return result
}

func Test_maxSubArray1(t *testing.T) {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := 6

	if maxSubArray1(arr) != expect {
		t.Fail()
	}

}
