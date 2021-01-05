package array

import (
	"fmt"
	"testing"
)

/**
两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	out := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out = append(out, i, j)
				return out
			}
		}
	}

	return out
}

func TestTwoSum(t *testing.T) {

	nums := []int{3, 2, 4}
	target := 6
	expect := []int{1, 2}

	out := twoSum(nums, target)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}

	nums = []int{2, 7, 11, 15}
	target = 9
	expect = []int{0, 1}

	out = twoSum(nums, target)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}

}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 2, 4}
		target := 6
		twoSum(nums, target)
	}
}

func twoSum1(nums []int, target int) []int {
	out := make([]int, 0)
	m := make(map[int]int)

	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			out = append(out, val, k)
		}
		m[v] = k
	}

	return out
}

func TestTwoSum1(t *testing.T) {

	nums := []int{3, 2, 4}
	target := 6
	expect := []int{1, 2}

	out := twoSum1(nums, target)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}

	nums = []int{2, 7, 11, 15}
	target = 9
	expect = []int{0, 1}

	out = twoSum1(nums, target)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}

}

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 2, 4}
		target := 6
		twoSum1(nums, target)
	}
}
