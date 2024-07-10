package resource

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
