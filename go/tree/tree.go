package main

import "errors"

// Tree .....
type Tree struct {
	Root *Node
}

// TreeValue ...
type TreeValue interface {
	Len() int
}

// Insert ....
func (t *Tree) Insert(value TreeValue) error {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return nil
	}
	return t.Root.Insert(value)
}

// Delete finds a treevalue and deletes it
func (t *Tree) Delete(s TreeValue) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	// Create a new node
	newParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, newParent)
	if err != nil {
		return err
	}

	// If we are the only node
	if newParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// Walk ...
func (t *Tree) Walk(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Walk(n.Left, f)
	f(n)
	t.Walk(n.Right, f)
}

func (t *Tree) WalkReverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.WalkReverse(n.Right, f)
	f(n)
	t.WalkReverse(n.Left, f)
}
