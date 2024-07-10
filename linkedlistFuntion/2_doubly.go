package linkedlistFuntion

import (
	"fmt"
)

// DoublyLinkedList의 리소스 선언
type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
}

// DoublyNode의 리소스 선언
type DoublyNode struct {
	// 노드가 담고있을 값
	DataField string
	// 다음 노드 주소
	LinkField *DoublyNode
	// 이전 노드 주소
	Prev *DoublyNode
}

// 1. 초기화된 노드 생성
func (l *DoublyLinkedList) NewDoublyNode(data string) {
	newNode := &DoublyNode{
		DataField: data,
		LinkField: nil,
		Prev:      nil,
	}
	l.Head = newNode
	l.Tail = newNode
}

// 2. Head 앞에 삽입
func (l *DoublyLinkedList) InsertDoublyFrontHead(data string) {
	// 새 노드를 선언하고 다음 노드를 Head로 설정
	newNode := &DoublyNode{
		DataField: data,
		LinkField: l.Head,
		Prev:      nil,
	}
	l.Head = newNode
}

// 3. Tail 뒤에 삽입
func (l *DoublyLinkedList) InsertDoublyBehindTail(data string) {
	// 새 노드를 선언하고 이전 노드를 Tail로 설정
	newNode := &DoublyNode{
		DataField: data,
		LinkField: nil,
		Prev:      l.Tail,
	}

	l.Tail = newNode

}

// 4. 특정 노드 값 뒤에 삽입
func (l *DoublyLinkedList) InsertDoublyBehindData(haveData, data string) {
	// haveData를 지닌 노드 탐색
	current := l.Head
	for current != nil {
		// 만약 현재 노드가 haveData를 지녔다면
		if current.DataField == haveData {
			// 새 노드를 선언하고 이전 노드를 current로 설정
			newNode := &DoublyNode{
				DataField: data,
				LinkField: nil,
				Prev:      current,
			}
			// 현재 노드의 다음 노드가 nil이라면 -> 현재노드가 Tail이라면
			if current.LinkField == nil {
				// 현재의 다음 노드는 새 노드
				current.LinkField = newNode
				l.Tail = newNode
				return
			}
			// 새 노드의 다음 노드는 현재의 다음 노드
			newNode.LinkField = current.LinkField
			// 현재 노드의 다음 노드의 이전 노드는 새 노드
			current.LinkField.Prev = newNode
			// 현재의 다음 노드는 새 노드
			current.LinkField = newNode
			return
		}
		// 계속 탐색
		current = current.LinkField

	}
	// 탐색이 끝났는데 haveData가 없다면 메세지 출력
	fmt.Printf("%s값을 지닌 노드가 없어 삽입 불가", haveData)
	fmt.Println()
}

// 5. 특정 노드 값 앞에 삽입
func (l *DoublyLinkedList) InsertDoublyFrontData(haveData, data string) {
	// Head가 haveData값을 지녔다면
	if l.Head.DataField == haveData {
		// 새 노드 선언
		newNode := &DoublyNode{
			DataField: data,
			LinkField: l.Head,
			Prev:      nil,
		}
		// Head에 newNode를 담기
		l.Head = newNode
		return
	}

	// haveData를 지닌 노드 탐색
	current := l.Head
	for current.LinkField != nil {
		// 현재 노드가 haveData를 지녔다면
		if current.DataField == haveData {
			// 다음 노드가 현재 노드인 새 노드 선언
			newNode := &DoublyNode{
				DataField: data,
				LinkField: current,
				Prev:      nil,
			}
			// 새 노드의 이전 노드는 현재 노드의 이전 노드
			newNode.Prev = current.LinkField
			// 현재 노드의 이전 노드의 다음 노드는 새 노드
			current.Prev.LinkField = newNode
			// 현재의 이전 노드는 newNode
			current.Prev = newNode
			return
		}
		// 계속 탐색
		current = current.LinkField
	}
	// 탐색이 끝났는데 haveData가 없다면 메세지 출력
	fmt.Printf("%s값을 지닌 노드가 없어 삽입 불가", haveData)
	fmt.Println()
}

// 1-1. 초기화된 노드 생성
func (l *DoublyLinkedList) NewDoublyCircularNode(data string) {
	newNode := &DoublyNode{
		DataField: data,
		LinkField: l.Head,
		Prev:      l.Head,
	}
	l.Head = newNode
	l.Tail = newNode
}

// 2-1. Head 앞에 삽입
func (l *DoublyLinkedList) InsertDoublyCircularFrontHead(data string) {
	// 새 노드를 선언하고 다음 노드를 Head로 설정
	newNode := &DoublyNode{
		DataField: data,
		LinkField: l.Head,
		Prev:      l.Tail,
	}
	// Head의 이전 노드는 새 노드
	l.Head.Prev = newNode
	// Head는 새 노드
	l.Head = newNode
	// Tail의 다음 노드를 newNode로 하여 원형 구현
	l.Tail.LinkField = newNode

}
