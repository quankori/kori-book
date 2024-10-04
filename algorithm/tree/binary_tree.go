package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
	} else {
		tree.Root.insert(value)
	}
}

func (node *Node) insert(value int) {
	if value < node.Value {
		// Left node
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.insert(value)
		}
	} else {
		// Right node
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.insert(value)
		}
	}
}

func (tree *BinaryTree) Search(value int) bool {
	return tree.Root.search(value)
}

func (node *Node) search(value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if value < node.Value {
		return node.Left.search(value)
	}
	return node.Right.search(value)
}

func (tree *BinaryTree) InOrderTraversal() []int {
	result := []int{}
	tree.Root.inOrderTraversal(&result)
	return result
}

func (node *Node) inOrderTraversal(result *[]int) {
	if node == nil {
		return
	}
	node.Left.inOrderTraversal(result)
	*result = append(*result, node.Value)
	node.Right.inOrderTraversal(result)
}

func (tree *BinaryTree) PreOrderTraversal() []int {
	result := []int{}
	tree.Root.preOrderTraversal(&result)
	return result
}

func (node *Node) preOrderTraversal(result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	node.Left.preOrderTraversal(result)
	node.Right.preOrderTraversal(result)
}

func (tree *BinaryTree) PostOrderTraversal() []int {
	result := []int{}
	tree.Root.postOrderTraversal(&result)
	return result
}

func (node *Node) postOrderTraversal(result *[]int) {
	if node == nil {
		return
	}
	node.Left.postOrderTraversal(result)
	node.Right.postOrderTraversal(result)
	*result = append(*result, node.Value)
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *BinaryTree) Delete(value int) {
	t.Root = deleteNode(t.Root, value)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	// Find value
	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Found root
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Value = minNode.Value
		root.Right = deleteNode(root.Right, minNode.Value)
	}
	return root
}
