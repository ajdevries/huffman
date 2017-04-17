package tree

import "testing"

var n = (&ID{}).Next

func TestNew(t *testing.T) {
	n := New("id", nil, nil)
	if n == nil {
		t.Fatal("Expecting a new Node")
	}
	if n.Left != nil || n.Right != nil {
		t.Fatal("Expecting no leafs")
	}
}

func TestNewLeaf(t *testing.T) {
	l := NewLeaf("id", "1", 1)
	if l == nil {
		t.Fatal("Expecting a new Leaf")
	}
	if l.Value != "1" {
		t.Fatal("Expecting value \"1\"")
	}
}

func TestHuffmanTwoValues(t *testing.T) {
	tree := Huffman("BA", n)
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.Left == nil || tree.Left.Value != "B" {
		t.Fatal("Expecting a Left leaf with value \"B\"")
	}
	if tree.Right == nil || tree.Right.Value != "A" {
		t.Fatal("Expecting a Right leaf with value \"A\"")
	}
}

func TestHuffmanThreeValues(t *testing.T) {
	tree := Huffman("CBA", n)
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.Left.Value != "A" {
		t.Fatal("Expecting a Left, Left leaf with value \"A\"")
	}
	if tree.Right.Left.Value != "C" {
		t.Fatal("Expecting a Left, Left leaf with value \"C\"")
	}
	if tree.Right.Right.Value != "B" {
		t.Fatal("Expecting a Right leaf with value \"B\"")
	}
}

func TestHuffmanFourValues(t *testing.T) {
	tree := Huffman("DCBA", n)
	if tree == nil {
		t.Fatal("Expecting a tree")
	}
	if tree.String() != "L:L:BR:AR:L:DR:C" {
		t.Fatal("Expecting tree value to be \"L:L:BR:AR:L:DR:C\"")
	}
}

func TestFind(t *testing.T) {
	tree := Huffman("AB", n)
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
	tree := Huffman("ABCDE", n)
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
	tree := Huffman("AB", n)
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
	tree := Huffman("ABCDE", n)
	_, err := tree.Decode("0")
	if err == nil {
		t.Fatal("Expecting an error")
	}
	v, _ := tree.Decode("110")
	if v != "E" {
		t.Fatalf("Expecting to decode 110 to value 'E' but got %s", v)
	}
}

func TestHuffmanDoubleValues(t *testing.T) {
	tree := Huffman("ABB", n)
	v, _ := tree.Decode("110")
	if v != "AAB" {
		t.Fatalf("Expecting to decode 110 to value 'ABB' but got %s", v)
	}
}

func TestEncode(t *testing.T) {
	tree := Huffman("ABB", n)
	v, _ := tree.Encode("ABB")
	if v != "100" {
		t.Fatalf("Expecting '100' but got %s", v)
	}
	_, err := tree.Encode("C")
	if err == nil {
		t.Fatal("Expecting an error")
	}
}

func TestNext(t *testing.T) {
	id := &ID{}
	if id.Next() != "1" {
		t.Fatal("Expecting 1")
	}
	if id.Next() != "2" {
		t.Fatal("Expecting 2")
	}
}
