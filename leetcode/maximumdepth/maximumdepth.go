package maximumdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := maxDepth(root.Right)
	l := maxDepth(root.Left)
	if r > l {
		return r + 1
	}
	return l + 1
}

/*
9 func maxDepth(root *TreeNode) int {
10    if root == nil{
11        return 0
12    }
13    return 1+ max(maxDepth(root.Right), maxDepth(root.Left))
14}
*/
