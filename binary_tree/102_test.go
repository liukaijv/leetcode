package binary_tree

import (
	"container/list"
	"fmt"
	"leetcode/common"
	"testing"
)

/**
二叉树的层序遍历

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
dfs
*/
func levelOrder(root *common.TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *common.TreeNode, level int, res [][]int) [][]int {

	if root == nil {
		return res
	}

	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}

	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)

	return res
}

func Test_levelOrder(t *testing.T) {

	var rootNode = &common.TreeNode{
		Val:  3,
		Left: &common.TreeNode{Val: 9},
		Right: &common.TreeNode{
			Val:   20,
			Left:  &common.TreeNode{Val: 15},
			Right: &common.TreeNode{Val: 7},
		},
	}

	expect := [][]int{{3}, {9, 20}, {15, 7}}

	out := levelOrder(rootNode)

	fmt.Println(out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}
}

/**
bfs
*/
func levelOrder1(root *common.TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushFront(root)

	deep := 0
	for queue.Len() > 0 {
		deep++
		var current []int
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Back()).(*common.TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, current)
	}
	fmt.Println(deep)
	return result
}

func Test_levelOrder1(t *testing.T) {

	var rootNode = &common.TreeNode{
		Val:  3,
		Left: &common.TreeNode{Val: 9},
		Right: &common.TreeNode{
			Val:   20,
			Left:  &common.TreeNode{Val: 15},
			Right: &common.TreeNode{Val: 7},
		},
	}

	expect := [][]int{{3}, {9, 20}, {15, 7}}

	out := levelOrder1(rootNode)

	fmt.Println(out)

	if fmt.Sprintf("%v", expect) != fmt.Sprintf("%v", out) {
		t.Fail()
	}
}
