package main

import (
	"fmt"
)

// LinkedList의 리소스 선언
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Node의 리소스 선언
type Node struct {
	// 노드가 담고있을 값
	DataField int
	// 다음 노드 주소
	LinkField *Node
}

func main() {
	list := LinkedList{}

	// 삽입 및 조회
	list.FirstNode(1)
	list.FindAll()
	list.InsertFrontHead(0)
	list.FindAll()
	list.InsertBehindTail(4)
	list.FindAll()
	list.InsertFrontData(4, 3)
	list.FindAll()
	list.InsertBehindData(1, 2)
	list.FindAll()
	list.FindHead()
	list.FindTail()
	list.CountNodes()

	// 삭제
	list.DeleteHead()
	list.FindAll()
	list.FindHead()
	list.DeleteTail()
	list.FindAll()
	list.FindTail()
	list.DeleteNode(3)
	list.FindAll()

}

// 1-1. 첫 노드 생성
func (l *LinkedList) FirstNode(data int) {
	newNode := &Node{
		DataField: data,
		LinkField: nil,
	}
	l.Head = newNode
	l.Tail = newNode
}

// 1-2. Head 앞에 삽입
func (l *LinkedList) InsertFrontHead(data int) {
	// 새 노드를 선언하고 다음 노드를 Head로 설정
	newNode := &Node{
		DataField: data,
		LinkField: l.Head,
	}
	l.Head = newNode
}

// 1-3. Tail 뒤에 삽입
func (l *LinkedList) InsertBehindTail(data int) {
	// 새 노드 선언
	newNode := &Node{
		DataField: data,
		LinkField: nil,
	}
	// Head부터 다음 노드가 없을 때 까지 마지막 노드를 탐색
	current := l.Head
	for current.LinkField != nil {
		current = current.LinkField
	}

	// 마지막 노드가 발견되면 다음 노드에 newNode를 담기
	current.LinkField = newNode
	l.Tail = newNode

}

// 1-4. 특정 노드 값 앞에 삽입
func (l *LinkedList) InsertFrontData(haveData, data int) {
	// Head가 haveData값을 지녔다면
	if l.Head.DataField == haveData {
		// 새 노드 선언
		newNode := &Node{
			DataField: data,
			LinkField: l.Head,
		}
		// Head에 newNode를 담기
		l.Head = newNode
		return
	}

	// haveData를 지닌 노드 탐색
	current := l.Head
	for current.LinkField != nil {
		// 현재 노드의 다음 노드가 haveData를 지녔다면
		if current.LinkField.DataField == haveData {
			// 다음 노드가 현재의 다음 노드인 새 노드 선언
			newNode := &Node{
				DataField: data,
				LinkField: current.LinkField,
			}
			// 현재의 다음 노드는 newNode
			current.LinkField = newNode
			return
		}
		// 계속 탐색
		current = current.LinkField
	}
	// 탐색이 끝났는데 haveData가 없다면 메세지 출력
	fmt.Printf("%d값을 지닌 노드가 없어 삽입 불가", haveData)
}

// 1-5. 특정 노드 값 뒤에 삽입
func (l *LinkedList) InsertBehindData(haveData, data int) {
	// haveData를 지닌 노드 탐색
	current := l.Head
	for current != nil {
		// 만약 현재 노드가 haveData를 지녔다면
		if current.DataField == haveData {
			// 새 노드 선언
			newNode := &Node{
				DataField: data,
				LinkField: nil,
			}
			// 현재 노드의 다음 노드가 nil이라면
			if current.LinkField == nil {
				// 현재의 다음 노드는 새 노드
				current.LinkField = newNode
				l.Tail = newNode
				return
			}
			// 새 노드의 다음 노드는 현재의 다음 노드
			newNode.LinkField = current.LinkField
			// 현재의 다음 노드는 새 노드
			current.LinkField = newNode
			return
		}
		// 계속 탐색
		current = current.LinkField

	}
	// 탐색이 끝났는데 haveData가 없다면 메세지 출력
	fmt.Printf("%d값을 지닌 노드가 없어 삽입 불가\n", haveData)
}

// 2-1. Head 탐색
func (l *LinkedList) FindHead() {
	current := l.Head
	if current != nil {
		fmt.Printf("Head : %d\n", current.DataField)
	}
}

// 2-2. Tail 탐색
func (l *LinkedList) FindTail() {
	current := l.Tail
	if current != nil {
		fmt.Printf("Tail : %d\n", current.DataField)
	}
}

// 2-3. 모든 값 탐색
func (l *LinkedList) FindAll() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ", current.DataField)
		current = current.LinkField
	}
	fmt.Println()
}

// 2-4. 크기 탐색
func (l *LinkedList) CountNodes() {
	current := l.Head
	var size int
	for current != nil {
		current = current.LinkField
		size++
	}
	fmt.Printf("길이 : %d\n", size)
}

// 3-1. Head 삭제
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

// 3-2. Tail 삭제
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

// 3-3. 특정 노드 삭제
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
			return
		}
		current = current.LinkField
	}
}
