package array

import (
	"fmt"
	"sort"
	"testing"
)

/**
三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
暴力遍历
再去重
*/
func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	l := len(nums)
	m := make(map[string]int)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					arr := []int{nums[i], nums[j], nums[k]}
					sort.Ints(arr)
					arrKey := fmt.Sprintf("%v", arr)
					if _, ok := m[arrKey]; !ok {
						out = append(out, arr)
					}
					m[fmt.Sprintf("%v", arr)] = 1
				}
			}
		}
	}

	return out
}

func Test_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	expect := [][]int{{-1, 0, 1}, {-1, -1, 2}}

	out := threeSum(nums)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}
}

func Benchmark_threeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{-1, 0, 1, 2, -1, -4}
		threeSum(nums)
	}
}

/**
双指针

*/
func threeSum1(nums []int) [][]int {
	out := make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l; i++ {
		//采取固定一个数，同时用双指针来查找另外两个数的方式
		target := 0 - nums[i]
		left := i + 1
		right := l - 1
		// 排好了序，如果固定下来的数（上面蓝色框框）本身就大于 0，那三数之和必然无法等于 0
		if nums[i] > 0 {
			break
		}

		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				// 如果和大于0，那就说明 right 的值太大，需要左移。如果和小于0，那就说明 left 的值太小，需要右移
				if nums[left]+nums[right] == target {
					arr := []int{nums[i], nums[left], nums[right]}
					out = append(out, arr)
					//left 和 right 是需要处理重复的情况，所以对于 left 和 left+1，以及 right 和 right-1，我们都单独做一下重复值的处理
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if (nums[left] + nums[right]) < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return out
}

func Test_threeSum1(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	expect := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	out := threeSum1(nums)

	fmt.Println(out)

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expect) {
		t.Fail()
	}
}

func Benchmark_threeSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{-1, 0, 1, 2, -1, -4}
		threeSum1(nums)
	}
}
