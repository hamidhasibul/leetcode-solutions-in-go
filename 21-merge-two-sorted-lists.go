package main

type ListNode struct {
	value int
	next  *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	temp := &ListNode{}
	current := temp

	if list1 != nil && list2 != nil {
		if list1.value < list2.value {
			current.next = list1
			list1 = list1.next
		} else {
			current.next = list2
			list2 = list2.next
		}
		current = current.next
	}

	if list1 != nil {
		current.next = list1
	}

	if list2 != nil {
		current.next = list2
	}

	return temp.next

}
