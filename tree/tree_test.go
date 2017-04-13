package tree

import "testing"

func TestNew(t *testing.T) {
	n := New(nil, nil)
	if n == nil {
		t.Fatal("Expecting a new Node")
	}
	if n.Left != nil || n.Right != nil {
		t.Fatal("Expecting no leafs")
	}
}

func TestNewLeaf(t *testing.T) {
	l := NewLeaf("1", 1)
	if l == nil {
		t.Fatal("Expecting a new Leaf")
	}
	if l.Value != "1" {
		t.Fatal("Expecting value \"1\"")
	}
}

func TestHuffmanTwoValues(t *testing.T) {
	tree := Huffman("B", "A")
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.Left == nil || tree.Left.Value != "A" {
		t.Fatal("Expecting a Left leaf with value \"A\"")
	}
	if tree.Right == nil || tree.Right.Value != "B" {
		t.Fatal("Expecting a Right leaf with value \"B\"")
	}
}

func TestHuffmanThreeValues(t *testing.T) {
	tree := Huffman("C", "B", "A")
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.Left == nil || tree.Left.Left.Value != "A" {
		t.Fatal("Expecting a Left, Left leaf with value \"A\"")
	}
	if tree.Left.Right.Value != "B" {
		t.Fatal("Expecting a Left, Left leaf with value \"B\"")
	}
	if tree.Right == nil || tree.Right.Value != "C" {
		t.Fatal("Expecting a Right leaf with value \"C\"")
	}
}

func TestHuffmanFourValues(t *testing.T) {
	tree := Huffman("D", "C", "B", "A")
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.Left.Left.Left.Value != "A" {
		t.Fatal("Expecting a Left, Left leaf with value \"A\"")
	}
	if tree.Left.Left.Right.Value != "B" {
		t.Fatal("Expecting a Left, Left leaf with value \"B\"")
	}
	if tree.Left.Right.Value != "C" {
		t.Fatal("Expecting a Right leaf with value \"C\"")
	}
	if tree.Right.Value != "D" {
		t.Fatal("Expecting a Right leaf with value \"D\"")
	}
	if tree.String() != "L:L:L:AR:BR:CR:D" {
		t.Fatal("Expecting tree value to be \"L:L:L:AR:BR:CR:D\"")
	}
}
