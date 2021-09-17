package main

import "errors"

// Node is the node in the tree
type Node struct {
	Left  *Node
	Right *Node
	Value TreeValue
}

// Equal is for comparing two treevalues
func (n *Node) Equal(a TreeValue) bool {
	return a == n.Value
}

// Greater is for sorting TreeValues
func (n *Node) Greater(a TreeValue) bool {
	return a.Len() > n.Value.Len()
}

// Less is for sorting TreeValue
func (n *Node) Less(a TreeValue) bool {
	return a.Len() < n.Value.Len()
}

// Insert .. inserts
func (n *Node) Insert(value TreeValue) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case n.Equal(value):
		return nil
	case n.Less(value):
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		}
		return n.Left.Insert(value)
	default:
		// If its not equal or less then it must be greater than
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return nil
		}
		return n.Right.Insert(value)
	}
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// Delete removes a node
func (n *Node) Delete(s TreeValue, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	// Search the node to be deleted.
	switch {
	case n.Less(s):
		return n.Left.Delete(s, n)
	case n.Greater(s):
		return n.Right.Delete(s, n)
	default:
		// If the node has no children remove it
		if n.Left == nil && n.Right == nil {
			return n.replaceNode(parent, nil)
			//			return nil
		}

		// If the node has one child: Replace the node with its child.
		if n.Left == nil {
			return n.replaceNode(parent, n.Right)
			//			return nil
		}
		if n.Right == nil {
			return n.replaceNode(parent, n.Left)
			//return nil
		}

		// If the node has two children:
		// Find the maximum element in the left subtree...
		replacement, replParent := n.Left.findMax(n)

		//...and replace the node's value and data with the replacement's value and data.
		n.Value = replacement.Value
		//		n.Data = replacement.Data

		// Then remove the replacement node.
		return replacement.Delete(replacement.Value, replParent)
	}
}
