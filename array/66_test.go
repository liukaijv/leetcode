package array

import (
	"fmt"
	"testing"
)

/**
加一

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。


示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]
 

提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func plusOne(digits []int) []int {

	if len(digits) < 1 || len(digits) > 100 {
		return digits
	}

	l := len(digits)
	carry := 0

	for i := l - 1; i >= 0; i-- {
		sum := 0
		// 加1
		if i == l-1 {
			sum = digits[i] + carry + 1
		} else {
			sum = digits[i] + carry
		}
		// 是否进位
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = sum % 10
	}

	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	fmt.Println(digits)

	return digits
}

func TestPlusOne(t *testing.T) {

	digits := []int{1, 2, 3}

	expect := []int{1, 2, 4}

	out := plusOne(digits)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

	digits = []int{9, 9, 9}

	expect = []int{1, 0, 0, 0}

	out = plusOne(digits)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

	digits = []int{1, 9, 9}

	expect = []int{2, 0, 0}

	out = plusOne(digits)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}

}
