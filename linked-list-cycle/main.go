func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    node1 := head
    node2 := node1.Next
    //a:=[]*ListNode{}
    for node1!=node2 {
        if node2 == nil || node2.Next == nil {
            return false
        }
        node1 = node1.Next
        node2 = node2.Next.Next
    }
    return true
}