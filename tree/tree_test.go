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
	if tree.Left.Value != "C" {
		t.Fatal("Expecting a Left, Left leaf with value \"A\"")
	}
	if tree.Right.Left.Value != "A" {
		t.Fatal("Expecting a Left, Left leaf with value \"B\"")
	}
	if tree.Right.Right.Value != "B" {
		t.Fatal("Expecting a Right leaf with value \"C\"")
	}
}

func TestHuffmanFourValues(t *testing.T) {
	tree := Huffman("D", "C", "B", "A")
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.String() != "L:L:CR:DR:L:AR:B" {
		t.Fatal("Expecting tree value to be \"L:L:CR:DR:L:AR:B\"")
	}
}

func TestFind(t *testing.T) {
	tree := Huffman("A", "B")
	f, _ := tree.Find("B")
	if f != "1" {
		t.Fatalf("Expecting 1 but got %s", f)
	}
	f, _ = tree.Find("A")
	if f != "0" {
		t.Fatalf("Expecting 0 but got %s", f)
	}
	_, err := tree.Find("Z")
	if err == nil {
		t.Fatal("Expecting an error")
	}
}

func TestFindSixValues(t *testing.T) {
	tree := Huffman("A", "B", "C", "D", "E")
	f, _ := tree.Find("B")
	if f != "000" {
		t.Fatalf("Expecting 1 but got %s", f)
	}
	f, _ = tree.Find("E")
	if f != "11" {
		t.Fatalf("Expecting 0 but got %s", f)
	}
}

func TestDecode(t *testing.T) {
	tree := Huffman("A", "B")
	v, _ := tree.Decode("0")
	if v != "A" {
		t.Fatalf("Expecting to decode 0 to value 'A' but got %s", v)
	}
	v, _ = tree.Decode("1")
	if v != "B" {
		t.Fatalf("Expecting to decode 1 to value 'B' but got %s", v)
	}
	v, _ = tree.Decode("110")
	if v != "BBA" {
		t.Fatalf("Expecting to decode 1 to value 'B' but got %s", v)
	}
}

func TestDecodeSixValues(t *testing.T) {
	tree := Huffman("A", "B", "C", "D", "E")
	_, err := tree.Decode("0")
	if err == nil {
		t.Fatal("Expecting an error")
	}
	v, _ := tree.Decode("110")
	if v != "E" {
		t.Fatalf("Expecting to decode 110 to value 'E' but got %s", v)
	}
}
