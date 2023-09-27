package pkg

// BSTNode is a node in a binary search tree
type BSTNode struct {
	Node  *Node
	Left  *BSTNode
	Right *BSTNode
}

// NewBSTNode creates a new BSTNode
func NewBSTNode(node *Node) *BSTNode {
	return &BSTNode{
		Node:  node,
		Left:  nil,
		Right: nil,
	}
}

// Insert inserts a node into the BST
func (bst *BSTNode) Insert(node *Node) {
	if node.hashedKey < bst.Node.hashedKey {
		if bst.Left == nil {
			bst.Left = NewBSTNode(node)
		} else {
			bst.Left.Insert(node)
		}
	} else {
		if bst.Right == nil {
			bst.Right = NewBSTNode(node)
		} else {
			bst.Right.Insert(node)
		}
	}
}

// Search searches for a node in the BST
func (bst *BSTNode) Search(hash uint64) *BSTNode {
	if bst == nil {
		return nil
	}
	if hash == bst.Node.hashedKey {
		return bst
	}
	if hash < bst.Node.hashedKey {
		return bst.Left.Search(hash)
	}
	return bst.Right.Search(hash)
}

// could be a weird edge case where we are searching for a hash that is too small, if
// too small then we should just return nil probably
func (bst *BSTNode) SearchLeftBiased(hash uint64, exact bool) *Node {
	closest := bst.Node
	for bst != nil {
		if exact && hash == bst.Node.hashedKey {
			return bst.Node
		}
		if hash < bst.Node.hashedKey {
			bst = bst.Left
		} else {
			closest = bst.Node
			bst = bst.Right
		}
	}
	return closest
}

func (bst *BSTNode) SearchRightBiased(hash uint64, exact bool) *Node {
	closest := bst.Node
	for bst != nil {
		if exact && hash == bst.Node.hashedKey {
			return bst.Node
		}
		if hash < bst.Node.hashedKey {
			closest = bst.Node
			bst = bst.Left
		} else {
			bst = bst.Right
		}
	}
	return closest
}

// Delete deletes a node from the BST
func (bst *BSTNode) Delete(hash uint64) *Node {
	// Delete looks for the node and finds successor to replace it

	// find the node
	if bst == nil {
		return nil
	}

	parent := bst

	for bst.Node.hashedKey != hash {
		if hash == bst.Node.hashedKey {
			break
		} else if hash < bst.Node.hashedKey {
			bst = bst.Left
		} else {
			bst = bst.Right
		}
		parent = bst
	}

	if bst.Left == nil && bst.Right == nil {
		// leaf node
		if parent.Left == bst {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if bst.Left == nil {
		// only right child
		if parent.Left == bst {
			parent.Left = bst.Right
		} else {
			parent.Right = bst.Right
		}
	} else if bst.Right == nil {
		// only left child
		if parent.Left == bst {
			parent.Left = bst.Left
		} else {
			parent.Right = bst.Left
		}
	} else {
		successor := bst.findSuccessor(bst)
		if successor == nil {
			return nil
		}
		bst.Node = successor.Node
	}

	return bst.Node
}

func (bst BSTNode) findSuccessor(node *BSTNode) *BSTNode {
	// find successor, which is the smallest node in the right subtree
	if node.Right == nil {
		return node
	}

	parent := node
	for node != nil {
		if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			break
		}
		parent = node
	}
	if parent.Left != nil {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
	return node
}

func (bst BSTNode) qsortAndBuild(nodes []Node) {
	quickSort(nodes, 0, len(nodes)-1)
	bst.build(nodes)
}

func (bst BSTNode) build(nodes []Node) {
	if nodes == nil || len(nodes) == 0 {
		return
	}
	mid := len(nodes) / 2
	bst.Insert(&nodes[mid])
	bst.build(nodes[:mid])
	bst.build(nodes[mid+1:])
}
