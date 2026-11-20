package mergelists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var start *ListNode
	if list1.Val <= list2.Val {
		start = list1
		list1 = list1.Next
	} else {
		start = list2
		list2 = list2.Next
	}
	target := start
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			target.Next = list1
			list1 = list1.Next
		} else {
			target.Next = list2
			list2 = list2.Next
		}
		target = target.Next
	}
	if list1 == nil {
		target.Next = list2
	}
	if list2 == nil {
		target.Next = list1
	}
	return start
}
