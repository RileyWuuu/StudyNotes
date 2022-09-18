package getNthElement

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l *LinkedList) getItem(n int) *Node {
	ptr := l.head
	if n < 0 {
		return ptr
	}
	if n > (l.length - 1) {
		return nil
	}
	for i := 0; i < n; i++ {
		ptr = ptr.next
	}
	return ptr
}
