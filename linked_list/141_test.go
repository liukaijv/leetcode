package linked_list

import (
	"leetcode/common"
	"testing"
)

/**
环形链表

给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。

进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

示例 1：

3->2->0->-4
  |	      |
   -------
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：

1->2
|  |
 --
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：

1
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：

链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
map实现
*/
func hasCycle(head *common.LinkedListNode) bool {

	if head == nil {
		return false
	}

	m := make(map[*common.LinkedListNode]int)

	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}
		m[head] = 1
		head = head.Next
	}

	return false
}

func Test_hasCycle(t *testing.T) {

	//head = [3,2,0,-4], pos = 1
	node1 := &common.LinkedListNode{Val: 3}
	node2 := &common.LinkedListNode{Val: 2}
	node3 := &common.LinkedListNode{Val: 0}
	node4 := &common.LinkedListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	if !hasCycle(node1) {
		t.Fail()
	}

}

/**
双指针
快指针每次多走一步
*/
func hasCycle1(head *common.LinkedListNode) bool {

	if head == nil {
		return false
	}

	fast := head.Next

	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}

	return false
}

func Test_hasCycle1(t *testing.T) {

	//head = [3,2,0,-4], pos = 1
	node1 := &common.LinkedListNode{Val: 3}
	node2 := &common.LinkedListNode{Val: 2}
	node3 := &common.LinkedListNode{Val: 0}
	node4 := &common.LinkedListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	if !hasCycle1(node1) {
		t.Fail()
	}

}
