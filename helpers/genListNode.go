package helpers

// ListNode define node
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// GenListNode 生成链表
func GenListNode(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	lNode := &ListNode{Value: values[0]}

	tmp := lNode
	for i := 1; i < len(values); i++ {
		tmp.Next = &ListNode{Value: values[i]}
		tmp = tmp.Next
	}

	return lNode
}
