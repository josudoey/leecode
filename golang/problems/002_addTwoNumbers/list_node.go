package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(values []int) *ListNode {
	var result *ListNode
	for i := len(values) - 1; i >= 0; i-- {
		result = &ListNode{
			Val:  values[i],
			Next: result,
		}
	}

	return result
}

func newIntSlice(node *ListNode) []int {
	var result []int

	for n := node; n != nil; n = n.Next {
		result = append(result, n.Val)
	}

	return result
}
