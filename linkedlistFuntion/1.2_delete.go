package linkedlistFuntion

// DeleteLinkedList의 리소스 선언
type DeleteLinkedList struct {
	Head *Node
	Tail *Node
}

// 1. Head 삭제
func (l *LinkedList) DeleteHead() {
	// Head 비어있을때
	if l.Head == nil {
		return

		// 리스트에 값이 한 개만 있을 때
	} else if l.Tail == l.Head {
		l.Head = nil
		l.Tail = nil

		// Head값을 Head값의 다음 노드
	} else {
		l.Head = l.Head.LinkField
	}
}

// 2. Tail 삭제
func (l *LinkedList) DeleteTail() {
	// Tail이 비어있을때
	if l.Tail == nil {
		return
	}

	// 리스트에 값이 한 개만 있을 때
	if l.Tail == l.Head {
		l.Head = nil
		l.Tail = nil
		return
	}

	// Tail의 뒷노드 탐색
	current := l.Head
	for current != nil {
		// 만약 현재 노드의 다음 노드가 Tail이면
		if current.LinkField == l.Tail {
			// 현재 노드의 다음 노드는 빈값
			current.LinkField = nil
			// Tail은 현재 노드
			l.Tail = current
		}
		current = current.LinkField
	}
}

// 3. 특정 노드 삭제
func (l *LinkedList) DeleteNode(haveData int) {
	// Head가 비어있을때
	if l.Head == nil {
		return
	}
	// 리스트에 값이 한 개만 있을 때
	if l.Tail == l.Head && l.Head.DataField == haveData {
		l.Head = nil
		l.Tail = nil
		return
	}

	// Head가 haveData를 지닐 때
	if l.Head.DataField == haveData {
		// Head는 다음 노드
		l.Head = l.Head.LinkField
		return
	}
	// haveData를 지닌 노드 탐색
	current := l.Head
	for current != nil {
		// 다음 노드의 데이터가 haveData를 지닐 때
		if current.LinkField.DataField == haveData {
			// 현재 노드의 다음 노드는 현재 노드의 다다음 노드
			current.LinkField = current.LinkField.LinkField
		}
		current = current.LinkField
	}
}
