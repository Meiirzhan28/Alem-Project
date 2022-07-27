package piscine

func ListReverse(l *List) {
	l.Tail = l.Head
	var prev *NodeL
	for l.Head != nil {
		temp := l.Head.Next
		l.Head.Next = prev
		prev = l.Head
		l.Head = temp
	}
	l.Head = prev
}
