package array

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := head.Next
	head.Next = swapPairs(newNode.Next)
	newNode.Next = head
	return newNode
}

func swapPairs2(head *ListNode) *ListNode {
	var newNode *ListNode
	for head != nil && head.Next != nil {

	}
	return newNode
}
