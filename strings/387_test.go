package strings

import "testing"

/**
字符串中的第一个唯一字符

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
桶
*/
func firstUniqChar(s string) int {
	var arr [26]int

	for _, v := range s {
		arr[v-'a']++
	}

	for k, v := range s {
		if arr[v-'a'] == 1 {
			return k
		}
	}

	return -1
}

/**
hash
*/
func firstUniqChar1(s string) int {

	m := make(map[int32]int)

	for _, v := range s {
		m[v]++
	}

	for k, v := range s {
		if m[v] == 1 {
			return k
		}
	}

	return -1
}

func Test_firstUniqChar(t *testing.T) {

	s := "loveleetcode"
	expect := 2

	if firstUniqChar(s) != expect {
		t.Fail()
	}
}

func Test_firstUniqChar1(t *testing.T) {

	s := "leetcode"
	expect := 0

	if firstUniqChar1(s) != expect {
		t.Fail()
	}
}
