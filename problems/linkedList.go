package problems

import (
	"fmt"
	"reflect"
)

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

func IsPalindrome(head *ListNode) bool {
	hv := []int{}
	hvt := []int{}
	tmp := head
	for i := 0; tmp != nil; i++ {
		hv = append(hv, tmp.Val)
		hvt = append(hvt, tmp.Val) // array or slice is pass by reference (you change one -swap here; other changes)
		tmp = tmp.Next
	}
	//swap array
	for i, j := 0, len(hv)-1; i < j; i, j = i+1, j-1 {
		hv[i], hv[j] = hv[j], hv[i]
	}
	if reflect.DeepEqual(hv, hvt) {
		return true
	}
	return false
}

func DeleteDuplicates(head *ListNode) *ListNode {
	result := head
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next //this will modify result for pass by reference
			// cause .Next this is memory location
		}
		head = head.Next //this will change head only not result cause this is pass by value
		// not dealing with a pointer here.
	}
	return result
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	if list1 == nil && list2 == nil {
		return nil
	} else if (list1 != nil && list2 != nil) && (list1.Val == 0 && list2.Val == 0) {
		return &ListNode{0, &ListNode{0, nil}}
	} else if (list1 != nil && list1.Val == 0) || (list2 != nil && list2.Val == 0) {
		return &ListNode{0, nil}
	}
	for list1 != nil || list2 != nil {
		if list1 != nil {
			tmp.Val = list1.Val
			list1 = list1.Next
		}
		if list2 != nil {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
			tmp.Val = list2.Val
			list2 = list2.Next
		}
		if list1 != nil || list2 != nil {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}
	return result
}
