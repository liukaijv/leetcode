package common

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) Add(val int) {
	newNode := &LinkedListNode{Val: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkedList) Remove(val int) {

}
