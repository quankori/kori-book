package linkedlist

import (
	"reflect"
	"testing"
)

func TestInsertAtHead(t *testing.T) {
	list := &SinglyLinkedList{}
	list.InsertAtHead(1)
	list.InsertAtHead(2)
	list.InsertAtHead(3)

	expected := []int{3, 2, 1}
	actual := list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAtTail(t *testing.T) {
	list := &SinglyLinkedList{}
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)

	expected := []int{1, 2, 3}
	actual := list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSearch(t *testing.T) {
	list := &SinglyLinkedList{}
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)

	if node := list.Search(2); node == nil || node.Value != 2 {
		t.Errorf("Expected to find 2, but did not")
	}

	if node := list.Search(4); node != nil {
		t.Errorf("Expected not to find 4, but did")
	}
}

func TestDelete(t *testing.T) {
	list := &SinglyLinkedList{}
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)

	list.Delete(2)
	expected := []int{1, 3}
	actual := list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	list.Delete(1)
	expected = []int{3}
	actual = list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	list.Delete(3)
	expected = []int{}
	actual = list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTraverse(t *testing.T) {
	list := &SinglyLinkedList{}
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)

	expected := []int{1, 2, 3}
	actual := list.Traverse()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
