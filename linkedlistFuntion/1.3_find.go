package linkedlistFuntion

import (
	"fmt"
)

type FindLinkedList struct {
	Head *Node
	Tail *Node
}

// 1. Head 탐색
func (l *FindLinkedList) FindHead() {
	current := l.Head
	if current != nil {
		fmt.Printf("Head : %d\n", current.DataField)
	}
}

// 2. Tail 탐색
func (l *FindLinkedList) FindTail() {
	current := l.Tail
	if current != nil {
		fmt.Printf("Tail : %d\n", current.DataField)
	}
}

// 3. 모든 값 출력
func (l *FindLinkedList) FindAll() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ", current.DataField)
		current = current.LinkField
	}
	fmt.Println()
}

// 4. 크기 탐색
func (l *FindLinkedList) CountNodes() (size int) {
	current := l.Head
	for current != nil {
		current = current.LinkField
		size++
	}
	return
}
