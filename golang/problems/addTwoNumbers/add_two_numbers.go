package addTwoNumbers

// ref https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var zeroNode = &ListNode{}

// branch tree
// zero carry
//  - l1,l2 not exist next
//  - only l1 exists next
//  - only l2 exists next
//  - l1,l2 exists next
// exists carry
//  - l1,l2 not exist next
//  - only l1 exists next
//  - only l2 exists next
//  - l1,l2 exists next

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	val := sum % 10   // remainder
	carry := sum / 10 // carry

	if carry == 0 {
		if l1.Next == nil && l2.Next == nil {
			return &ListNode{Val: val}
		} else if l2.Next == nil {
			return &ListNode{
				Val:  val,
				Next: addTwoNumbers(l1.Next, zeroNode),
			}
		} else if l1.Next == nil {
			return &ListNode{
				Val:  val,
				Next: addTwoNumbers(l2.Next, zeroNode),
			}
		}

		return &ListNode{
			Val:  val,
			Next: addTwoNumbers(l1.Next, l2.Next),
		}
	}

	carryNode := &ListNode{Val: carry}
	if l1.Next == nil && l2.Next == nil {
		return &ListNode{
			Val:  val,
			Next: carryNode,
		}
	} else if l2.Next == nil {
		return &ListNode{
			Val:  val,
			Next: addTwoNumbers(l1.Next, carryNode),
		}
	} else if l1.Next == nil {
		return &ListNode{
			Val:  val,
			Next: addTwoNumbers(l2.Next, carryNode),
		}
	}

	return &ListNode{
		Val: val,
		Next: addTwoNumbers(
			addTwoNumbers(l1.Next, l2.Next),
			carryNode,
		),
	}
}
