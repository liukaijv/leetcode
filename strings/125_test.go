package strings

import (
	"regexp"
	"strings"
	"testing"
)

/**
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
回文串就是对称的字符串，如level
双指针
*/
func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^a-z0-9]")

	s = reg.ReplaceAllString(s, "")

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func Test_isPalindrome(t *testing.T) {

	s := "A man, a plan, a canal: Panama"

	if !isPalindrome(s) {
		t.Fail()
	}
}

func Test_isPalindrome1(t *testing.T) {

	s := "race a car"

	if isPalindrome(s) {
		t.Fail()
	}

	if !isPalindrome("level") {
		t.Fail()
	}
}
