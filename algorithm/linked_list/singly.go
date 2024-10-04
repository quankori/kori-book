package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type SinglyLinkedList struct {
	Head *Node
}

func (list *SinglyLinkedList) InsertAtHead(value int) {
	newNode := &Node{Value: value}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *SinglyLinkedList) InsertAtTail(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *SinglyLinkedList) Search(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *SinglyLinkedList) Delete(value int) {
	if list.Head == nil {
		return
	}
	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}
	current := list.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (list *SinglyLinkedList) Traverse() []int {
	result := []int{}
	current := list.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
