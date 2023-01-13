package problems

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) {
	result := &ListNode{}
	tmp := result // address like 0xc000010240 of result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val //Val goes on to variable at address which is result
			l1 = l1.Next      //l1 is diminished to nil throughout the for loop
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next //l2 is diminished to nil throughout the for loop
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{} // Next was created for variable at tmp address which is 'result' and for tmp
		}
		tmp = tmp.Next // this step will make tmp move to nil eventually, not variable result
	}
	fmt.Println(result)
}
