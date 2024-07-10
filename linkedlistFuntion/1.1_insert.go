package linkedlistFuntion

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

// 1. 첫 노드 생성
func (l *LinkedList) FirstNode(data int) {
	newNode := &Node{
		DataField: data,
		LinkField: nil,
	}
	l.Head = newNode
	l.Tail = newNode
}

// 2. Head 앞에 삽입
func (l *LinkedList) InsertFrontHead(data int) {
	// 새 노드를 선언하고 다음 노드를 Head로 설정
	newNode := &Node{
		DataField: data,
		LinkField: l.Head,
	}
	l.Head = newNode
}

// 3. Tail 뒤에 삽입
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

// 4. 특정 노드 값 앞에 삽입
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

// 5. 특정 노드 값 뒤에 삽입
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
	fmt.Printf("%d값을 지닌 노드가 없어 삽입 불가", haveData)
}

// 1-1. 초기화된 원형 노드 생성
func (l *LinkedList) NewCircularNode(data int) {
	newNode := &Node{
		DataField: data,
		LinkField: l.Head,
	}
	l.Head = newNode
	l.Tail = newNode
}

// 2-1. Head 앞에 삽입
func (l *LinkedList) InsertCircularFrontHead(data int) {
	// 새 노드를 선언하고 다음 노드를 Head로 설정
	newNode := &Node{
		DataField: data,
		LinkField: l.Head,
	}
	l.Head = newNode
	// Tail의 다음 노드를 newNode로 하여 원형 구현
	l.Tail.LinkField = newNode

}
