package tree

import (
	"reflect"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	tree := &BinaryTree{}
	values := []int{5, 3, 7, 2, 4, 6, 8}

	for _, value := range values {
		tree.Insert(value)
	}

	for _, value := range values {
		if !tree.Search(value) {
			t.Errorf("Expected to find %d in the tree", value)
		}
	}

	if tree.Search(10) {
		t.Errorf("Expected not to find 10 in the tree")
	}
}

func TestInOrderTraversal(t *testing.T) {
	tree := &BinaryTree{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, value := range values {
		tree.Insert(value)
	}

	expectedOrder := []int{2, 3, 4, 5, 6, 7, 8}
	actualOrder := tree.InOrderTraversal()
	if !reflect.DeepEqual(actualOrder, expectedOrder) {
		t.Errorf("Expected in-order traversal to be %v, got %v", expectedOrder, actualOrder)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := &BinaryTree{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, value := range values {
		tree.Insert(value)
	}

	expectedOrder := []int{5, 3, 2, 4, 7, 6, 8}
	actualOrder := tree.PreOrderTraversal()
	if !reflect.DeepEqual(actualOrder, expectedOrder) {
		t.Errorf("Expected pre-order traversal to be %v, got %v", expectedOrder, actualOrder)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := &BinaryTree{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, value := range values {
		tree.Insert(value)
	}

	expectedOrder := []int{2, 4, 3, 6, 8, 7, 5}
	actualOrder := tree.PostOrderTraversal()
	if !reflect.DeepEqual(actualOrder, expectedOrder) {
		t.Errorf("Expected pre-order traversal to be %v, got %v", expectedOrder, actualOrder)
	}
}

func TestDelete(t *testing.T) {
	tree := &BinaryTree{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, value := range values {
		tree.Insert(value)
	}

	tree.Delete(3)
	expectedOrderAfterDeletion := []int{2, 4, 5, 6, 7, 8}
	actualOrderAfterDeletion := tree.InOrderTraversal()
	if !reflect.DeepEqual(actualOrderAfterDeletion, expectedOrderAfterDeletion) {
		t.Errorf("Expected in-order traversal after deletion to be %v, got %v", expectedOrderAfterDeletion, actualOrderAfterDeletion)
	}

	if tree.Search(3) {
		t.Errorf("Expected not to find 3 in the tree after deletion")
	}
}
