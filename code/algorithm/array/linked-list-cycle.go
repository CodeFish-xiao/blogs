package array

// 遍历方式
func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return ok
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	fast := head.Next
	slow := head.Next.Next
	for fast != slow {
		if slow == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next.Next
	}
	return true
}
