package linkedlist

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (l *LinkedList) Init() {
	l.Head = nil
	l.Size = 0
}

func (l *LinkedList) Insert(x int) {
	node := Node{
		Val:  x,
		Next: nil,
	}
	if l.Empty() {
		l.Head = &node
	} else {
		pre := l.Head
		for pre.Next != nil {
			pre = pre.Next
		}
		pre.Next = &node
	}
	l.Size++
}

func (l *LinkedList) Delete(x int) {
	if l.Empty() {
		return
	}
	pre := &Node{Next: l.Head}
	for pre.Next != nil {
		if pre.Next.Val == x {
			if pre.Next == l.Head {
				l.Head = l.Head.Next
			} else {
				pre.Next = pre.Next.Next
			}
			l.Size--
			break
		}
		pre = pre.Next
	}
}

func (l *LinkedList) Empty() bool {
	return l.Size == 0
}

func (l *LinkedList) Length() int {
	return l.Size
}
