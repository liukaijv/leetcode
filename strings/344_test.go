package strings

import (
	"strings"
	"testing"
)

/**
反转字符串

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：

输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
双指针
*/
func reverseString(s []string) []string {

	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}

func reverseString2(s []string) []string {

	l := len(s)

	for i := 0; i < len(s)/2; i++ {
		l--
		s[i], s[l] = s[l], s[i]
	}

	return s
}

func Test_reverseString(t *testing.T) {

	s := []string{"h", "e", "l", "l", "o"}
	expect := []string{"o", "l", "l", "e", "h"}

	out := reverseString(s)

	if strings.Join(out, "") != strings.Join(expect, "") {
		t.Fail()
	}
}

func Test_reverseString2(t *testing.T) {

	s := []string{"h", "e", "l", "l", "o"}
	expect := []string{"o", "l", "l", "e", "h"}

	out := reverseString2(s)

	if strings.Join(out, "") != strings.Join(expect, "") {
		t.Fail()
	}
}
