package balancedtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return checkHight(root) != -1
}
func checkHight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	r := checkHight(node.Right)
	if r == -1 {
		return -1
	}
	l := checkHight(node.Left)
	if l == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return 1 + max(r, l)
}
