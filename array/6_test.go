package array

import (
	"fmt"
	"strings"
	"testing"
)

/**
字形变换

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
找规律

0     6      12      18
1   5 7   11 13   17 19
2 4   8 10   14 16   20
3     9      15      21

于n行的, s中的第i个字符：

对余数进行判断

i%(2n-2) == 0 ----> row0

i%(2n-2) == 1 & 2n-2-1 ----> row1

i%(2n-2) == 2 & 2n-2-2 ----> row2

...

i%(2n-2) == n-1 ----> row(n-1)

==>

是以2n-2周期
对 k = i%(2n-2)进行判断

k<=n-1时候，s[i]就属于第k行
k>n-1时候，s[i]就属于2n-2-k行
最后将rows拼接起来就行了
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var b = []rune(s)
	fmt.Println(b)
	var res = make([]string, numRows)
	var length = len(b)
	var period = numRows*2 - 2
	for i := 0; i < length; i++ {
		var k = i % period
		if k < numRows {
			res[k] += string(b[i])
		} else {
			res[period-k] += string(b[i])
		}
	}
	return strings.Join(res, "")
}

func Test_convert(t *testing.T) {

	s, numRows := "LEETCODEISHIRING", 3
	expect := "LCIRETOESIIGEDHN"

	out := convert(s, numRows)
	fmt.Println(out)

	if out != expect {
		t.Fail()
	}
}
